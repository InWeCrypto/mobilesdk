package ethmobile

import (
	"encoding/hex"
	"math/big"
	"strings"

	"github.com/inwecrypto/bip39"
	"github.com/inwecrypto/ethgo"
	"github.com/inwecrypto/ethgo/erc20"
	"github.com/inwecrypto/ethgo/keystore"
	"github.com/inwecrypto/ethgo/tx"
)

// Wallet neo mobile wallet
type Wallet struct {
	key *keystore.Key
}

// New create a new wallet
func New() (*Wallet, error) {
	key, err := keystore.NewKey()

	if err != nil {
		return nil, err
	}

	return &Wallet{
		key: key,
	}, nil
}

// FromPrivateKey .
func FromPrivateKey(privateKey []byte) (*Wallet, error) {
	key, err := keystore.KeyFromPrivateKey(privateKey)

	if err != nil {
		return nil, err
	}

	return &Wallet{
		key: key,
	}, nil
}

// FromMnemonic create wallet from mnemonic
func FromMnemonic(mnemonic string, lang string) (*Wallet, error) {
	dic, _ := bip39.GetDict(lang)

	data, err := bip39.MnemonicToByteArray(mnemonic, dic)

	if err != nil {
		return nil, err
	}

	data = data[1 : len(data)-1]

	println(hex.EncodeToString(data))

	key, err := keystore.KeyFromPrivateKey(data)

	if err != nil {
		return nil, err
	}

	return &Wallet{
		key: key,
	}, nil
}

// FromKeyStore create wallet from keystore
func FromKeyStore(ks string, password string) (*Wallet, error) {
	key, err := keystore.ReadKeyStore([]byte(ks), password)

	if err != nil {
		return nil, err
	}

	return &Wallet{
		key: key,
	}, nil
}

// ToKeyStore write wallet to keystore format string
func (wallet *Wallet) ToKeyStore(password string) (string, error) {
	keystore, err := keystore.WriteLightScryptKeyStore(wallet.key, password)

	return string(keystore), err
}

// Transfer transfer eth to target address
func (wallet *Wallet) Transfer(nonce, to, amount, gasPrice, gasLimits string) (string, error) {

	nonceBigInt, err := readBigint(nonce)

	if err != nil {
		return "", err
	}

	amountBigInt, err := readBigint(amount)

	if err != nil {
		return "", err
	}

	gasPriceBigInt, err := readBigint(gasPrice)

	if err != nil {
		return "", err
	}

	gasLimitsBigInt, err := readBigint(gasLimits)

	if err != nil {
		return "", err
	}

	rawTx := tx.NewTx(
		nonceBigInt.Uint64(),
		to,
		(*ethgo.Value)(amountBigInt),
		(*ethgo.Value)(gasPriceBigInt),
		gasLimitsBigInt,
		nil)

	data, err := rawTx.Encode()

	if err != nil {
		return "", err
	}

	return hex.EncodeToString(data), nil
}

// TransferERC20 transfer eth to target address
func (wallet *Wallet) TransferERC20(contract, nonce, to, amount, gasPrice, gasLimits string) (string, error) {

	codes, err := erc20.Transfer(to, amount)

	if err != nil {
		return "", err
	}

	nonceBigInt, err := readBigint(nonce)

	if err != nil {
		return "", err
	}

	gasPriceBigInt, err := readBigint(gasPrice)

	if err != nil {
		return "", err
	}

	gasLimitsBigInt, err := readBigint(gasLimits)

	if err != nil {
		return "", err
	}

	rawTx := tx.NewTx(
		nonceBigInt.Uint64(),
		contract,
		nil,
		(*ethgo.Value)(gasPriceBigInt),
		gasLimitsBigInt,
		codes)

	data, err := rawTx.Encode()

	if err != nil {
		return "", err
	}

	return hex.EncodeToString(data), nil
}

func readBigint(source string) (*big.Int, error) {
	value := big.NewInt(0)

	if source == "0x0" {
		return value, nil
	}

	source = strings.TrimPrefix(source, "0x")

	if len(source)%2 != 0 {
		source = "0" + source
	}

	data, err := hex.DecodeString(source)

	if err != nil {
		return nil, err
	}

	return value.SetBytes(data), nil
}
