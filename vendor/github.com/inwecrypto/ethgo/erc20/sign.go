package erc20

import (
	"encoding/hex"
	"fmt"
	"strings"

	"github.com/inwecrypto/sha3"
)

const (
	signBalanceOf         = "balanceOf(address)"
	signTotalSupply       = "totalSupply()"
	signTransfer          = "transfer(address,uint256)"
	signTransferFrom      = "transferFrom(address,address,uint256)"
	signApprove           = "approve(address,uint256)"
	signName              = "name()"
	signSymbol            = "symbol()"
	signAllowance         = "allowance(address,address)"
	eventTransfer         = "Transfer(address,address,uint256)"
	decimals              = "decimals()"
	signTransferOwnership = "transferOwnership(address)"
)

// Method/Event id
var (
	TransferID          = SignABI(signTransfer)
	BalanceOfID         = SignABI(signBalanceOf)
	Decimals            = SignABI(decimals)
	TransferFromID      = SignABI(signTransferFrom)
	ApproveID           = SignABI(signApprove)
	TotalSupplyID       = SignABI(signTotalSupply)
	AllowanceID         = SignABI(signAllowance)
	TransferOwnershipID = SignABI(signTransferOwnership)
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
	address = packNumeric(address, 32)

	return fmt.Sprintf("0x%s%s", BalanceOfID, address)
}

// GetDecimals .
func GetDecimals() string {
	return fmt.Sprintf("0x%s", Decimals)
}

func GetTotalSupply() string {
	return fmt.Sprintf("0x%s", TotalSupplyID)
}

func GetName() string {
	return "0x" + SignABI(signName)
}

func GetSignSymbol() string {
	return "0x" + SignABI(signSymbol)
}

func packNumeric(value string, bytes int) string {
	if value == "" {
		value = "0x0"
	}

	value = strings.TrimPrefix(value, "0x")

	chars := bytes * 2

	n := len(value)
	if n%chars == 0 {
		return value
	}
	return strings.Repeat("0", chars-n%chars) + value
}

// Transfer .
func Transfer(to string, value string) ([]byte, error) {
	to = packNumeric(to, 32)
	value = packNumeric(value, 32)

	data := fmt.Sprintf("%s%s%s", SignABI(signTransfer), to, value)

	return hex.DecodeString(data)
}

// TransferFrom .
func TransferFrom(from, to string, value string) ([]byte, error) {
	from = packNumeric(from, 32)
	to = packNumeric(to, 32)
	value = packNumeric(value, 32)

	data := fmt.Sprintf("%s%s%s%s", TransferFromID, from, to, value)

	return hex.DecodeString(data)
}

// Approve .
func Approve(to string, value string) ([]byte, error) {
	to = packNumeric(to, 32)
	value = packNumeric(value, 32)

	data := fmt.Sprintf("%s%s%s", ApproveID, to, value)

	return hex.DecodeString(data)
}

func Allowance(from, to string) ([]byte, error) {
	from = packNumeric(from, 32)
	to = packNumeric(to, 32)

	data := fmt.Sprintf("%s%s%s", AllowanceID, to, to)

	return hex.DecodeString(data)
}

func TransferOwnership(to string) ([]byte, error) {
	to = packNumeric(to, 32)
	data := fmt.Sprintf("%s%s", TransferOwnershipID, to)

	return hex.DecodeString(data)
}
