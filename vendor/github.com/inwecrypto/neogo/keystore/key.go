package keystore

import (
	"bytes"
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/sha256"
	"fmt"
	"math/big"
	"strings"

	"github.com/inwecrypto/keystore"
	"github.com/pborman/uuid"
	"golang.org/x/crypto/ripemd160"
)

// const variables
var (
	StandardScryptN = 1 << 18
	StandardScryptP = 1
	LightScryptN    = 1 << 12
	LightScryptP    = 6
)

// Key wallet wallet key
type Key struct {
	ID         uuid.UUID         // Key ID
	Address    string            // address
	PrivateKey *ecdsa.PrivateKey // btc private key
}

// NewKey create new key
func NewKey() (*Key, error) {
	privateKey, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)

	if err != nil {
		return nil, err
	}

	return &Key{
		ID:         uuid.NewRandom(),
		PrivateKey: privateKey,
		Address:    toNeoAddress(&privateKey.PublicKey),
	}, nil
}

func toECDSA(key []byte, curve elliptic.Curve) *ecdsa.PrivateKey {
	priv := new(ecdsa.PrivateKey)
	priv.PublicKey.Curve = curve
	priv.D = new(big.Int).SetBytes(key)
	priv.PublicKey.X, priv.PublicKey.Y = priv.PublicKey.Curve.ScalarBaseMult(key)
	return priv
}

// FromWIF converts a Wallet Import Format string to a Bitcoin private key and derives the corresponding Bitcoin public key.
func FromWIF(wif string, curve elliptic.Curve) (*ecdsa.PrivateKey, error) {
	/* See https://en.bitcoin.it/wiki/Wallet_import_format */

	/* Base58 Check Decode the WIF string */
	ver, privBytes, err := b58checkdecode(wif)
	if err != nil {
		return nil, err
	}

	/* Check that the version byte is 0x80 */
	if ver != 0x80 {
		return nil, fmt.Errorf("Invalid WIF version 0x%02x, expected 0x80", ver)
	}

	/* If the private key bytes length is 33, check that suffix byte is 0x01 (for compression) and strip it off */
	if len(privBytes) == 33 {
		if privBytes[len(privBytes)-1] != 0x01 {
			return nil, fmt.Errorf("Invalid private key, unknown suffix byte 0x%02x", privBytes[len(privBytes)-1])
		}
		privBytes = privBytes[0:32]
	}

	return toECDSA(privBytes, curve), nil
}

// KeyFromPrivateKey wallet key from private key bytes
func KeyFromPrivateKey(privateKeyBytes []byte) (*Key, error) {
	privateKey := toECDSA(privateKeyBytes, elliptic.P256())

	return &Key{
		ID:         uuid.NewRandom(),
		PrivateKey: privateKey,
		Address:    toNeoAddress(&privateKey.PublicKey),
	}, nil
}

// KeyFromWIF wallet key from wif format
func KeyFromWIF(wif string) (*Key, error) {
	privateKey, err := FromWIF(wif, elliptic.P256())

	if err != nil {
		return nil, err
	}

	return &Key{
		ID:         uuid.NewRandom(),
		PrivateKey: privateKey,
		Address:    toNeoAddress(&privateKey.PublicKey),
	}, nil
}

func keystoreKeyToNEOKey(key *keystore.Key) (*Key, error) {

	privateKey := toECDSA(key.PrivateKey, elliptic.P256())

	return &Key{
		ID:         uuid.UUID(key.ID),
		Address:    key.Address,
		PrivateKey: privateKey,
	}, nil
}

func toBytes(priv *ecdsa.PrivateKey) (b []byte) {
	d := priv.D.Bytes()

	/* Pad D to 32 bytes */
	paddedd := append(bytes.Repeat([]byte{0x00}, 32-len(d)), d...)

	return paddedd
}

func neoKeyToKeyStoreKey(key *Key) (*keystore.Key, error) {
	bytes := toBytes(key.PrivateKey)

	return &keystore.Key{
		ID:         key.ID,
		Address:    key.Address,
		PrivateKey: bytes,
	}, nil
}

// ToBytes get key's bytes array
func (key *Key) ToBytes() []byte {
	return toBytes(key.PrivateKey)
}

// WriteScryptKeyStore write keystore with Scrypt format
func WriteScryptKeyStore(key *Key, password string) ([]byte, error) {
	keyStoreKey, err := neoKeyToKeyStoreKey(key)

	if err != nil {
		return nil, err
	}

	attrs := map[string]interface{}{
		"ScryptN": StandardScryptN,
		"ScryptP": StandardScryptP,
	}

	return keystore.Encrypt(keyStoreKey, password, attrs)
}

// WriteLightScryptKeyStore write keystore with Scrypt format
func WriteLightScryptKeyStore(key *Key, password string) ([]byte, error) {
	keyStoreKey, err := neoKeyToKeyStoreKey(key)

	if err != nil {
		return nil, err
	}

	attrs := map[string]interface{}{
		"ScryptN": LightScryptN,
		"ScryptP": LightScryptP,
	}

	return keystore.Encrypt(keyStoreKey, password, attrs)
}

// ReadKeyStore read key from keystore
func ReadKeyStore(data []byte, password string) (*Key, error) {
	keystore, err := keystore.Decrypt(data, password)

	if err != nil {
		return nil, err
	}

	return keystoreKeyToNEOKey(keystore)
}

func toNeoAddress(publickKey *ecdsa.PublicKey) (address string) {
	/* See https://en.bitcoin.it/wiki/Technical_background_of_Bitcoin_addresses */

	x := publickKey.X.Bytes()

	/* Pad X to 32-bytes */
	paddedx := append(bytes.Repeat([]byte{0x00}, 32-len(x)), x...)

	var pubbytes []byte

	/* Add prefix 0x02 or 0x03 depending on ylsb */
	if publickKey.Y.Bit(0) == 0 {
		pubbytes = append([]byte{0x02}, paddedx...)
	} else {
		pubbytes = append([]byte{0x03}, paddedx...)
	}

	pubbytes = append([]byte{0x21}, pubbytes...)
	pubbytes = append(pubbytes, 0xAC)

	/* SHA256 Hash */
	sha256h := sha256.New()
	sha256h.Reset()
	sha256h.Write(pubbytes)
	pubhash1 := sha256h.Sum(nil)

	/* RIPEMD-160 Hash */
	ripemd160h := ripemd160.New()
	ripemd160h.Reset()
	ripemd160h.Write(pubhash1)
	pubhash2 := ripemd160h.Sum(nil)

	programhash := pubhash2

	//wallet version
	//program_hash = append([]byte{0x17}, program_hash...)

	// doublesha := sha256Bytes(sha256Bytes(program_hash))

	// checksum := doublesha[0:4]

	// result := append(program_hash, checksum...)
	/* Convert hash bytes to base58 check encoded sequence */
	address = b58checkencodeNEO(0x17, programhash)

	return address
}

func b58checkencodeNEO(ver uint8, b []byte) (s string) {
	/* Prepend version */
	bcpy := append([]byte{ver}, b...)

	/* Create a new SHA256 context */
	sha256h := sha256.New()

	/* SHA256 Hash #1 */
	sha256h.Reset()
	sha256h.Write(bcpy)
	hash1 := sha256h.Sum(nil)

	/* SHA256 Hash #2 */
	sha256h.Reset()
	sha256h.Write(hash1)
	hash2 := sha256h.Sum(nil)

	/* Append first four bytes of hash */
	bcpy = append(bcpy, hash2[0:4]...)

	/* Encode base58 string */
	s = b58encode(bcpy)

	// /* For number of leading 0's in bytes, prepend 1 */
	// for _, v := range bcpy {
	// 	if v != 0 {
	// 		break
	// 	}
	// 	s = "1" + s
	// }

	return s
}

func b58encode(b []byte) (s string) {
	/* See https://en.bitcoin.it/wiki/Base58Check_encoding */

	const BitcoinBase58Table = "123456789ABCDEFGHJKLMNPQRSTUVWXYZabcdefghijkmnopqrstuvwxyz"

	/* Convert big endian bytes to big int */
	x := new(big.Int).SetBytes(b)

	/* Initialize */
	r := new(big.Int)
	m := big.NewInt(58)
	zero := big.NewInt(0)
	s = ""

	/* Convert big int to string */
	for x.Cmp(zero) > 0 {
		/* x, r = (x / 58, x % 58) */
		x.QuoRem(x, m, r)
		/* Prepend ASCII character */
		s = string(BitcoinBase58Table[r.Int64()]) + s
	}

	return s
}

// b58decode decodes a base-58 encoded string into a byte slice b.
func b58decode(s string) (b []byte, err error) {
	/* See https://en.bitcoin.it/wiki/Base58Check_encoding */

	const BitcoinBase58Table = "123456789ABCDEFGHJKLMNPQRSTUVWXYZabcdefghijkmnopqrstuvwxyz"

	/* Initialize */
	x := big.NewInt(0)
	m := big.NewInt(58)

	/* Convert string to big int */
	for i := 0; i < len(s); i++ {
		b58index := strings.IndexByte(BitcoinBase58Table, s[i])
		if b58index == -1 {
			return nil, fmt.Errorf("Invalid base-58 character encountered: '%c', index %d", s[i], i)
		}
		b58value := big.NewInt(int64(b58index))
		x.Mul(x, m)
		x.Add(x, b58value)
	}

	/* Convert big int to big endian bytes */
	b = x.Bytes()

	return b, nil
}

/******************************************************************************/
/* Base-58 Check Encode/Decode */
/******************************************************************************/

// b58checkencode encodes version ver and byte slice b into a base-58 check encoded string.
func b58checkencode(ver uint8, b []byte) (s string) {
	/* Prepend version */
	bcpy := append([]byte{ver}, b...)

	/* Create a new SHA256 context */
	sha256h := sha256.New()

	/* SHA256 Hash #1 */
	sha256h.Reset()
	sha256h.Write(bcpy)
	hash1 := sha256h.Sum(nil)

	/* SHA256 Hash #2 */
	sha256h.Reset()
	sha256h.Write(hash1)
	hash2 := sha256h.Sum(nil)

	/* Append first four bytes of hash */
	bcpy = append(bcpy, hash2[0:4]...)

	/* Encode base58 string */
	s = b58encode(bcpy)

	/* For number of leading 0's in bytes, prepend 1 */
	for _, v := range bcpy {
		if v != 0 {
			break
		}
		s = "1" + s
	}

	return s
}

// b58checkdecode decodes base-58 check encoded string s into a version ver and byte slice b.
func b58checkdecode(s string) (ver uint8, b []byte, err error) {
	/* Decode base58 string */
	b, err = b58decode(s)
	if err != nil {
		return 0, nil, err
	}

	/* Add leading zero bytes */
	for i := 0; i < len(s); i++ {
		if s[i] != '1' {
			break
		}
		b = append([]byte{0x00}, b...)
	}

	/* Verify checksum */
	if len(b) < 5 {
		return 0, nil, fmt.Errorf("Invalid base-58 check string: missing checksum")
	}

	/* Create a new SHA256 context */
	sha256h := sha256.New()

	/* SHA256 Hash #1 */
	sha256h.Reset()
	sha256h.Write(b[:len(b)-4])
	hash1 := sha256h.Sum(nil)

	/* SHA256 Hash #2 */
	sha256h.Reset()
	sha256h.Write(hash1)
	hash2 := sha256h.Sum(nil)

	/* Compare checksum */
	if bytes.Compare(hash2[0:4], b[len(b)-4:]) != 0 {
		return 0, nil, fmt.Errorf("invalid base-58 check string: invalid checksum")
	}

	/* Strip checksum bytes */
	b = b[:len(b)-4]

	/* Extract and strip version */
	ver = b[0]
	b = b[1:]

	return ver, b, nil
}
