package nep5

import (
	"bytes"
	"math/big"

	"github.com/inwecrypto/neogo"
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
	writer := neogo.NewScriptWriter(&buff)

	writer.
		EmitPushInteger(amount).
		EmitPushBytes(to).
		EmitPushBytes(from).
		EmitPushInteger(big.NewInt(3)).
		Emit(neogo.PACK, nil).
		EmitPushString("transfer").
		EmitAPPCall(scriptHash, false)

	if nil != writer.Error {
		return nil, writer.Error
	}

	return buff.Bytes(), nil
}

// MintToken .
func MintToken(scriptHash []byte) ([]byte, error) {
	var buff bytes.Buffer
	writer := neogo.NewScriptWriter(&buff)

	writer.
		EmitPushInteger(big.NewInt(0)).
		Emit(neogo.PACK, nil).
		EmitPushString("mintTokens").
		EmitAPPCall(scriptHash, false)

	if nil != writer.Error {
		return nil, writer.Error
	}

	return buff.Bytes(), nil
}
