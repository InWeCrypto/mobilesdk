package tx

import (
	"crypto/ecdsa"
	"encoding/hex"
	"math/big"
	"strings"

	"github.com/inwecrypto/ethgo"
	"github.com/inwecrypto/ethgo/math"
	"github.com/inwecrypto/ethgo/rlp"
	"github.com/inwecrypto/gosecp256k1"
	"github.com/inwecrypto/sha3"
)

// const asset .
const (
	EthAsset = "0x0000000000000000000000000000000000000000"
)

// Tx .
type Tx struct {
	AccountNonce uint64    `json:"nonce"    gencodec:"required"`
	Price        *big.Int  `json:"gasPrice" gencodec:"required"`
	GasLimit     *big.Int  `json:"gas"      gencodec:"required"`
	Recipient    *[20]byte `json:"to"       rlp:"nil"` // nil means contract creation
	Amount       *big.Int  `json:"value"    gencodec:"required"`
	Payload      []byte    `json:"input"    gencodec:"required"`
	V            *big.Int  `json:"v" gencodec:"required"`
	R            *big.Int  `json:"r" gencodec:"required"`
	S            *big.Int  `json:"s" gencodec:"required"`
}

// NewTx create new eth tx
func NewTx(nonce uint64, to string, amount, gasPrice *ethgo.Value, gasLimit *big.Int, data []byte) *Tx {

	to = strings.TrimPrefix(to, "0x")

	var recipient [20]byte

	toBytes, _ := hex.DecodeString(to)

	copy(recipient[:], toBytes)

	tx := &Tx{
		AccountNonce: nonce,
		Recipient:    &recipient,
		Payload:      data,
		Amount:       (*big.Int)(amount),
		GasLimit:     (*big.Int)(gasLimit),
		Price:        (*big.Int)(gasPrice),
		V:            new(big.Int),
		R:            new(big.Int),
		S:            new(big.Int),
	}

	return tx
}

// Sign .
func (tx *Tx) Sign(prv *ecdsa.PrivateKey) error {
	hw := sha3.NewKeccak256()

	rlp.Encode(hw, []interface{}{
		tx.AccountNonce,
		tx.Price,
		tx.GasLimit,
		tx.Recipient,
		tx.Amount,
		tx.Payload,
	})

	var hash [32]byte

	hw.Sum(hash[:0])

	seckey := math.PaddedBigBytes(prv.D, prv.Params().BitSize/8)

	sig, err := secp256k1.Sign(hash[:], seckey)

	if err != nil {
		return err
	}

	zeroBytes(seckey)

	tx.R = new(big.Int).SetBytes(sig[:32])
	tx.S = new(big.Int).SetBytes(sig[32:64])
	tx.V = new(big.Int).SetBytes([]byte{sig[64] + 27})

	return nil
}

// Encode .
func (tx *Tx) Encode() ([]byte, error) {
	return rlp.EncodeToBytes(tx)
}

func zeroBytes(bytes []byte) {
	for i := range bytes {
		bytes[i] = 0
	}
}
