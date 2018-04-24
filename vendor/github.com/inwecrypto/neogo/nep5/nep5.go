package nep5

import (
	"bytes"
	"math/big"

	"github.com/inwecrypto/neogo/script"
)

// ParameterType .
type ParameterType byte

// Parameter Type enum
const (
	Signature ParameterType = 0
	Boolean
	Integer
	Hash160
	Hash256
	ByteArray
	PublicKey
	String
	Array
	InteropInterface
	Void
)

// Contract neo nep5 contract object
type Contract struct {
	scriptHash []byte
}

// NewContract .
func NewContract(scriptHash []byte) *Contract {
	return &Contract{
		scriptHash: scriptHash,
	}
}

// Transfer implement nep5 transfer method
// more detail visit website https://github.com/neo-project/proposals/blob/master/nep-5.mediawiki#trasfer
func Transfer(scriptHash []byte, from []byte, to []byte, amount *big.Int) ([]byte, error) {
	var buff bytes.Buffer
	transferScript := script.New("transfer")
	// writer := neogo.NewScriptWriter(&buff)

	transferScript.
		EmitPushInteger(amount).
		EmitPushBytes(to).
		EmitPushBytes(from).
		EmitPushInteger(big.NewInt(3)).
		Emit(script.PACK, nil).
		EmitPushString("transfer").
		EmitAPPCall(scriptHash, false)

	err := transferScript.Write(&buff)

	if err != nil {
		return nil, err
	}

	return buff.Bytes(), nil
}

// MintToken .
func MintToken(scriptHash []byte) ([]byte, error) {
	var buff bytes.Buffer
	writer := script.New("mint")

	writer.
		EmitPushInteger(big.NewInt(0)).
		Emit(script.PACK, nil).
		EmitPushString("mintTokens").
		EmitAPPCall(scriptHash, false)

	if err := writer.Write(&buff); err != nil {
		return nil, err
	}

	return buff.Bytes(), nil
}

// DeployContract .
func DeployContract(script []byte, parmeters []ParameterType) ([]byte, error) {
	return nil, nil
}
