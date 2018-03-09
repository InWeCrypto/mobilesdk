package keystore

import (
	"bytes"
	"crypto/ecdsa"
	"crypto/elliptic"
	"fmt"

	"github.com/btcsuite/btcutil/base58"
	"github.com/inwecrypto/neogo/script"
)

// DecodeWIF .
func DecodeWIF(wif string) (*ecdsa.PrivateKey, error) {
	bytesOfPrivateKey, version, err := base58.CheckDecode(wif)

	if err != nil {
		return nil, err
	}

	/* Check that the version byte is 0x80 */
	if version != 0x80 {
		return nil, fmt.Errorf("Invalid WIF version 0x%02x, expected 0x80", version)
	}

	/* If the private key bytes length is 33, check that suffix byte is 0x01 (for compression) and strip it off */
	if len(bytesOfPrivateKey) == 33 {
		if bytesOfPrivateKey[len(bytesOfPrivateKey)-1] != 0x01 {
			return nil, fmt.Errorf("Invalid private key, unknown suffix byte 0x%02x", bytesOfPrivateKey[len(bytesOfPrivateKey)-1])
		}
		bytesOfPrivateKey = bytesOfPrivateKey[0:32]
	}

	return toECDSA(bytesOfPrivateKey, elliptic.P256()), nil
}

// EncodeWIF .
func EncodeWIF(privateKey *ecdsa.PrivateKey) (string, error) {
	bytesOfPrivateKey := privateKey.D.Bytes()

	if len(bytesOfPrivateKey) == 32 {
		bytesOfPrivateKey = append(bytesOfPrivateKey, 0x01)
	}

	return base58.CheckEncode(bytesOfPrivateKey, 0x80), nil
}

// PrivateToScriptHash .
func PrivateToScriptHash(privateKey *ecdsa.PrivateKey) ([]byte, error) {
	publicKey := privateKey.PublicKey

	x := publicKey.X.Bytes()

	/* Pad X to 32-bytes */
	paddedx := append(bytes.Repeat([]byte{0x00}, 32-len(x)), x...)

	var pubbytes []byte

	/* Add prefix 0x02 or 0x03 depending on ylsb */
	if publicKey.Y.Bit(0) == 0 {
		pubbytes = append([]byte{0x02}, paddedx...)
	} else {
		pubbytes = append([]byte{0x03}, paddedx...)
	}

	addressScript := script.New("address")

	addressScript.EmitPushBytes(pubbytes)
	addressScript.Emit(script.CHECKSIG, nil)

	return addressScript.Hash()
}

// PrivateToAddress .
func PrivateToAddress(privateKey *ecdsa.PrivateKey) (string, error) {

	programhash, err := PrivateToScriptHash(privateKey)

	if err != nil {
		return "", err
	}

	return base58.CheckEncode(programhash, 0x17), nil
}
