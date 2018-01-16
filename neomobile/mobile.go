package neomobile

import (
	"encoding/hex"
	"encoding/json"
	"math/big"
	"strings"

	"github.com/inwecrypto/bip39"
	"github.com/inwecrypto/neogo/keystore"
	"github.com/inwecrypto/neogo/nep5"
	"github.com/inwecrypto/neogo/rpc"
	neotx "github.com/inwecrypto/neogo/tx"
)

// Wallet neo mobile wallet
type Wallet struct {
	key *keystore.Key
}

// Tx neo rawtx wrapper
type Tx struct {
	Data string
	ID   string
}

// FromWIF create wallet from wif
func FromWIF(wif string) (*Wallet, error) {
	key, err := keystore.KeyFromWIF(wif)

	if err != nil {
		return nil, err
	}

	return &Wallet{
		key: key,
	}, nil
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
func (wrapper *Wallet) ToKeyStore(password string) (string, error) {
	keystore, err := keystore.WriteLightScryptKeyStore(wrapper.key, password)

	return string(keystore), err
}

// CreateAssertTx create assert transfer raw tx
func (wrapper *Wallet) CreateAssertTx(assert, from, to string, amount float64, unspent string) (*Tx, error) {
	var utxos []*rpc.UTXO

	if err := json.Unmarshal([]byte(unspent), &utxos); err != nil {
		return nil, err
	}

	vout := []*neotx.Vout{
		&neotx.Vout{
			Asset:   assert,
			Value:   neotx.MakeFixed8(amount),
			Address: to,
		},
	}

	tx := neotx.NewContractTx()

	err := tx.CalcInputs(vout, utxos)

	if err != nil {
		return nil, err
	}

	// tx, err := neotx.CreateSendAssertTx(assert, from, to, amount, utxos)

	// if err != nil {
	// 	return nil, err
	// }

	rawtxdata, txid, err := tx.Tx().Sign(wrapper.key.PrivateKey)

	return &Tx{
		Data: hex.EncodeToString(rawtxdata),
		ID:   txid,
	}, err
}

// Address get wallet address
func (wrapper *Wallet) Address() string {
	return wrapper.key.Address
}

// Mnemonic gete mnemonic string
func (wrapper *Wallet) Mnemonic(lang string) (string, error) {
	privateKeyBytes := wrapper.key.ToBytes()

	dic, _ := bip39.GetDict(lang)

	println(hex.EncodeToString(privateKeyBytes))

	data, err := bip39.NewMnemonic(privateKeyBytes, dic)

	if err != nil {
		return "", err
	}

	return data, nil
}

// CreateClaimTx create claim tx
func (wrapper *Wallet) CreateClaimTx(amount float64, address string, unspent string) (*Tx, error) {
	var utxos []*rpc.UTXO

	if err := json.Unmarshal([]byte(unspent), &utxos); err != nil {
		return nil, err
	}

	tx := neotx.NewClaimTx()

	err := tx.Claim(amount, address, utxos)

	if err != nil {
		return nil, err
	}

	rawtxdata, txid, err := tx.Tx().Sign(wrapper.key.PrivateKey)

	return &Tx{
		Data: hex.EncodeToString(rawtxdata),
		ID:   txid,
	}, err
}

// MintToken .
func (wrapper *Wallet) MintToken(asset string, gas, amount float64, unspent string) (*Tx, error) {
	var utxos []*rpc.UTXO

	if err := json.Unmarshal([]byte(unspent), &utxos); err != nil {
		return nil, err
	}

	scriptHash, err := hex.DecodeString(strings.TrimPrefix(asset, "0x"))

	if err != nil {
		return nil, err
	}

	scriptHash = reverseBytes(scriptHash)

	address := neotx.EncodeAddress(scriptHash)

	script, err := nep5.MintToken(scriptHash)

	if err != nil {
		return nil, err
	}

	tx := neotx.NewInvocationTx(script, gas)

	vout := []*neotx.Vout{
		&neotx.Vout{
			Asset:   neotx.NEOAssert,
			Value:   neotx.MakeFixed8(amount),
			Address: address,
		},
	}

	err = tx.CalcInputs(vout, utxos)

	if err != nil {
		return nil, err
	}

	rawtxdata, txid, err := tx.Tx().Sign(wrapper.key.PrivateKey)

	return &Tx{
		Data: hex.EncodeToString(rawtxdata),
		ID:   txid,
	}, err
}

// CreateNep5Tx create nep5 transfer transaction
func (wrapper *Wallet) CreateNep5Tx(asset string, from, to string, gas float64, amount int64, unspent string) (*Tx, error) {

	var utxos []*rpc.UTXO

	if err := json.Unmarshal([]byte(unspent), &utxos); err != nil {
		return nil, err
	}

	scriptHash, err := hex.DecodeString(strings.TrimPrefix(asset, "0x"))

	if err != nil {
		return nil, err
	}

	scriptHash = reverseBytes(scriptHash)

	bytesOfFrom, err := hex.DecodeString(from)

	if err != nil {
		return nil, err
	}

	bytesOfFrom = reverseBytes(bytesOfFrom)

	bytesOfTo, err := hex.DecodeString(to)

	if err != nil {
		return nil, err
	}

	bytesOfTo = reverseBytes(bytesOfTo)

	script, err := nep5.Transfer(scriptHash, bytesOfFrom, bytesOfTo, big.NewInt(amount))

	tx := neotx.NewInvocationTx(script, gas)

	err = tx.CalcInputs(nil, utxos)

	if err != nil {
		return nil, err
	}

	rawtxdata, txid, err := tx.Tx().Sign(wrapper.key.PrivateKey)

	return &Tx{
		Data: hex.EncodeToString(rawtxdata),
		ID:   txid,
	}, err
}

func reverseBytes(s []byte) []byte {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}

	return s
}

// DecodeAddress decode address
func DecodeAddress(address string) (string, error) {
	bytesOfAddress, err := neotx.DecodeAddress(address)

	if err != nil {
		return "", err
	}

	bytesOfAddress = reverseBytes(bytesOfAddress)

	return hex.EncodeToString(bytesOfAddress), nil
}

// EncodeAddress encode address
func EncodeAddress(address string) (string, error) {

	bytesOfAddress, err := hex.DecodeString(strings.TrimPrefix(address, "0x"))

	if err != nil {
		return "", err
	}

	bytesOfAddress = reverseBytes(bytesOfAddress)

	return neotx.EncodeAddress(bytesOfAddress), nil
}
