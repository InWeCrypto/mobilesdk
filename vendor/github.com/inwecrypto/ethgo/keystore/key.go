package keystore

import (
	"bytes"
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"encoding/hex"
	"math/big"
	"strings"

	"github.com/inwecrypto/gosecp256k1"
	"github.com/inwecrypto/sha3"

	"github.com/inwecrypto/keystore"
	"github.com/pborman/uuid"
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
	privateKey, err := ecdsa.GenerateKey(secp256k1.S256(), rand.Reader)

	if err != nil {
		return nil, err
	}

	return &Key{
		ID:         uuid.NewRandom(),
		PrivateKey: privateKey,
		Address:    pubKeyToAddress(&privateKey.PublicKey),
	}, nil
}

func toECDSA(key []byte, curve elliptic.Curve) *ecdsa.PrivateKey {
	priv := new(ecdsa.PrivateKey)
	priv.PublicKey.Curve = curve
	priv.D = new(big.Int).SetBytes(key)
	priv.PublicKey.X, priv.PublicKey.Y = priv.PublicKey.Curve.ScalarBaseMult(key)
	return priv
}

// KeyFromPrivateKey wallet key from private key bytes
func KeyFromPrivateKey(privateKeyBytes []byte) (*Key, error) {
	privateKey := toECDSA(privateKeyBytes, secp256k1.S256())

	return &Key{
		ID:         uuid.NewRandom(),
		PrivateKey: privateKey,
		Address:    pubKeyToAddress(&privateKey.PublicKey),
	}, nil
}

func keystoreKeyToNEOKey(key *keystore.Key) (*Key, error) {

	privateKey := toECDSA(key.PrivateKey, elliptic.P256())

	address := key.Address

	if !strings.HasPrefix(address, "0x") {
		address = "0x" + address
	}

	return &Key{
		ID:         uuid.UUID(key.ID),
		Address:    address,
		PrivateKey: privateKey,
	}, nil
}

func pubKeyToAddress(pub *ecdsa.PublicKey) string {
	pubBytes := pubKeyBytes(pub)

	hasher := sha3.NewKeccak256()

	hasher.Write(pubBytes[1:])

	pubBytes = hasher.Sum(nil)[12:]

	if len(pubBytes) > 20 {
		pubBytes = pubBytes[len(pubBytes)-20:]
	}

	address := make([]byte, 20)

	copy(address[20-len(pubBytes):], pubBytes)

	unchecksummed := hex.EncodeToString(address)

	sha := sha3.NewKeccak256()

	sha.Write([]byte(unchecksummed))

	hash := sha.Sum(nil)

	result := []byte(unchecksummed)

	for i := 0; i < len(result); i++ {
		hashByte := hash[i/2]
		if i%2 == 0 {
			hashByte = hashByte >> 4
		} else {
			hashByte &= 0xf
		}
		if result[i] > '9' && hashByte > 7 {
			result[i] -= 32
		}
	}

	return "0x" + string(result)
}

func pubKeyBytes(pub *ecdsa.PublicKey) []byte {
	if pub == nil || pub.X == nil || pub.Y == nil {
		return nil
	}
	return elliptic.Marshal(secp256k1.S256(), pub.X, pub.Y)
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
