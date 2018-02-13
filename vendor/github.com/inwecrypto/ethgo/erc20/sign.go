package erc20

import (
	"encoding/hex"
	"fmt"
	"strings"

	"github.com/inwecrypto/sha3"
)

const (
	signBalanceOf   = "balanceOf(address)"
	signTotalSupply = "totalSupply()"
	signTransfer    = "transfer(address,uint256)"
	eventTransfer   = "Transfer(address,address,uint256)"
	initWallet      = "initWallet(address[],uint256,uint256)"
	decimals        = "decimals()"
)

// Method/Event id
var (
	TransferID   = SignABI(signTransfer)
	BalanceOfID  = SignABI(signBalanceOf)
	InitWalletID = SignABI(initWallet)
	Decimals     = SignABI(decimals)
)

// SignABI sign abi string
func SignABI(abi string) string {
	hasher := sha3.NewKeccak256()
	hasher.Write([]byte(abi))
	data := hasher.Sum(nil)

	return hex.EncodeToString(data[0:4])
}

// BalanceOf create erc20 balanceof abi string
func BalanceOf(address string) string {
	address = strings.Trim(address, "0x")

	return fmt.Sprintf("0x%s%s", BalanceOfID, packNumeric(address, 32))
}

// GetDecimals .
func GetDecimals() string {
	return fmt.Sprintf("0x%s", Decimals)
}

func packNumeric(value string, bytes int) string {
	value = strings.TrimSuffix(value, "0x")

	chars := bytes * 2

	n := len(value)
	if n%chars == 0 {
		return value
	}
	return strings.Repeat("0", chars-n%chars) + value
}

// Transfer .
func Transfer(to string, value string) ([]byte, error) {
	to = packNumeric(strings.TrimPrefix(to, "0x"), 32)
	value = packNumeric(strings.TrimPrefix(value, "0x"), 32)

	data := fmt.Sprintf("%s%s%s", SignABI(signTransfer), to, value)
	println(data)

	return hex.DecodeString(data)
}
