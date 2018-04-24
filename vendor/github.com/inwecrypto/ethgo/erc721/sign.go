package erc721

import (
	"encoding/hex"
	"fmt"
	"math/big"
	"strings"

	"github.com/inwecrypto/sha3"
)

const (
	ERC721_balanceOf           = "balanceOf(address)"
	ERC721_totalSupply         = "totalSupply()"
	ERC721_transfer            = "transfer(address,uint256)"
	ERC721_decimals            = "decimals()"
	ERC721_name                = "name()"
	ERC721_symbol              = "symbol()"
	ERC721_ownerOf             = "ownerOf(uint256)"
	ERC721_approve             = "approve(address,uint256)"
	ERC721_setApprovalForAll   = "setApprovalForAll(address,bool)"
	ERC721_getApprovedAddress  = "getApprovedAddress(uint256)"
	ERC721_isApprovedForAll    = "isApprovedForAll(address,address)"
	ERC721_takeOwnership       = "takeOwnership(uint256)"
	ERC721_tokenOfOwnerByIndex = "tokenOfOwnerByIndex(address,uint256)"
	ERC721_tokenMetadata       = "tokenMetadata(uint256)"
	ERC721_tokensOf            = "tokensOf(address)"
	ERC721_exists              = "exists(uint256)"
	ERC721_setAssetHolder      = "setAssetHolder(address,uint256)"
	ERC721_transferFrom        = "transferFrom(address,address,uint256)"
	ERC721_isAuthorized        = "isAuthorized(address,uint256)"
	ERC721_description         = "description()"

	DecentraLand_decodeTokenId = "decodeTokenId(uint256)"
	DecentraLand_encodeTokenId = "encodeTokenId(int256,int256)"
	DecentraLand_landData      = "landData(int256,int256)"
	DecentraLand_landOf        = "landOf(address)"
	DecentraLand_transferLand  = "transferLand(int,int,address)"
	DecentraLand_ownerOfLand   = "ownerOfLand(int,int)"

	RedPacket_newRedPacket       = "newRedPacket(address,address,uint256,uint256,uint256)"
	RedPacket_open               = "open(uint256,address,uint256)"
	RedPacket_openMany           = "openMany(uint256,address[],uint256)"
	RedPacket_takeBack           = "takeBack(uint256)"
	RedPacket_setTaxCost         = "setTaxCost(uint256)"
	RedPacket_changeWallet       = "changeWallet(address)"
	RedPacket_changeMaxCount     = "changeMaxCount(uint256)"
	RedPacket_getRedPacketDetail = "getRedPacketDetail(uint256)"
	RedPacket_sendEther          = "sendEther(uint256)"
	RedPacket_taxCost            = "taxCost()"
)

// Method/Event id
var (
	Method_transfer            = SignABI(ERC721_transfer)
	Method_balanceOf           = SignABI(ERC721_balanceOf)
	Method_decimals            = SignABI(ERC721_decimals)
	Method_totalSupply         = SignABI(ERC721_totalSupply)
	Method_name                = SignABI(ERC721_name)
	Method_symbol              = SignABI(ERC721_symbol)
	Method_ownerOf             = SignABI(ERC721_ownerOf)
	Method_approve             = SignABI(ERC721_approve)
	Method_setApprovalForAll   = SignABI(ERC721_setApprovalForAll)
	Method_getApprovedAddress  = SignABI(ERC721_getApprovedAddress)
	Method_isApprovedForAll    = SignABI(ERC721_isApprovedForAll)
	Method_takeOwnership       = SignABI(ERC721_takeOwnership)
	Method_tokenOfOwnerByIndex = SignABI(ERC721_tokenOfOwnerByIndex)
	Method_tokenMetadata       = SignABI(ERC721_tokenMetadata)
	Method_tokensOf            = SignABI(ERC721_tokensOf)
	Method_exists              = SignABI(ERC721_exists)
	Method_setAssetHolder      = SignABI(ERC721_setAssetHolder)
	Method_transferFrom        = SignABI(ERC721_transferFrom)
	Method_isAuthorized        = SignABI(ERC721_isAuthorized)
	Method_description         = SignABI(ERC721_description)

	Method_DecentraLand_decodeTokenId = SignABI(DecentraLand_decodeTokenId)
	Method_DecentraLand_encodeTokenId = SignABI(DecentraLand_encodeTokenId)
	Method_DecentraLand_landData      = SignABI(DecentraLand_landData)
	Method_DecentraLand_landOf        = SignABI(DecentraLand_landOf)
	Method_DecentraLand_transferLand  = SignABI(DecentraLand_transferLand)
	Method_DecentraLand_ownerOfLand   = SignABI(DecentraLand_ownerOfLand)
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

	return fmt.Sprintf("0x%s%s", Method_balanceOf, address)
}

// GetDecimals .
func GetDecimals() string {
	return fmt.Sprintf("0x%s", Method_decimals)
}

func GetName() string {
	return fmt.Sprintf("0x%s", Method_name)
}

func GetSymbol() string {
	return fmt.Sprintf("0x%s", Method_symbol)
}

func GetDescription() string {
	return fmt.Sprintf("0x%s", Method_description)
}

func OwnerOf(value string) string {
	value = packNumeric(value, 32)
	return fmt.Sprintf("0x%s%s", Method_ownerOf, value)
}

func TokensOf(address string) string {
	address = packNumeric(address, 32)

	return fmt.Sprintf("0x%s%s", Method_tokensOf, address)
}

func SetAssetHolder(to string, value string) ([]byte, error) {
	to = packNumeric(to, 32)
	value = packNumeric(value, 32)

	data := fmt.Sprintf("%s%s%s", Method_setAssetHolder, to, value)

	return hex.DecodeString(data)
}

func GetTokenMetadata(value string) string {
	value = packNumeric(value, 32)
	return fmt.Sprintf("0x%s%s", Method_tokenMetadata, value)
}

func Approve(to string, value string) ([]byte, error) {
	to = packNumeric(to, 32)
	value = packNumeric(value, 32)

	data := fmt.Sprintf("%s%s%s", Method_approve, to, value)
	return hex.DecodeString(data)
}

func packNumeric(value string, bytes int) string {
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

	data := fmt.Sprintf("%s%s%s", Method_transfer, to, value)

	return hex.DecodeString(data)
}

func TransferLand(to string, x, y string) ([]byte, error) {
	to = packNumeric(to, 32)
	x = packNumeric(x, 32)
	y = packNumeric(y, 32)

	data := fmt.Sprintf("%s%s%s%s", Method_DecentraLand_transferLand, x, y, to)

	return hex.DecodeString(data)
}

func TransferFrom(from, to string, value string) ([]byte, error) {
	from = packNumeric(from, 32)
	to = packNumeric(to, 32)
	value = packNumeric(value, 32)

	data := fmt.Sprintf("%s%s%s%s", Method_transferFrom, from, to, value)

	return hex.DecodeString(data)
}

func IsExists(value string) string {
	value = packNumeric(value, 32)

	return fmt.Sprintf("0x%s%s", Method_exists, value)
}

func TokenOfOwnerByIndex(adress string, value string) string {
	adress = packNumeric(adress, 32)
	value = packNumeric(value, 32)

	return fmt.Sprintf("0x%s%s%s", Method_tokenOfOwnerByIndex, adress, value)
}

func TakeOwnership(value string) ([]byte, error) {
	data := fmt.Sprintf("%s%s", Method_takeOwnership, value)

	return hex.DecodeString(data)
}

func DecodeTokenId(value string) string {
	value = packNumeric(value, 32)

	return fmt.Sprintf("0x%s%s", Method_DecentraLand_decodeTokenId, value)
}

func EncodeTokenId(x, y string) string {
	x = packNumeric(x, 32)
	y = packNumeric(y, 32)

	return fmt.Sprintf("0x%s%s%s", Method_DecentraLand_encodeTokenId, x, y)
}

func LandData(x, y string) string {
	x = packNumeric(x, 32)
	y = packNumeric(y, 32)

	return fmt.Sprintf("0x%s%s%s", Method_DecentraLand_landData, x, y)
}

func Description() string {
	return fmt.Sprintf("0x%s", Method_description)
}

func LandOf(address string) string {
	address = packNumeric(address, 32)

	return fmt.Sprintf("0x%s%s", Method_DecentraLand_landOf, address)
}

func OwnerOfLand(x, y string) string {
	x = packNumeric(x, 32)
	y = packNumeric(y, 32)

	return fmt.Sprintf("0x%s%s%s", Method_DecentraLand_ownerOfLand, x, y)
}

func TaxCost() ([]byte, error) {
	data := SignABI(RedPacket_taxCost)

	return hex.DecodeString(data)
}

func SetTaxCost(value string) ([]byte, error) {
	data := SignABI(RedPacket_setTaxCost) + packNumeric(value, 32)

	return hex.DecodeString(data)
}

func ChangeWallet(address string) ([]byte, error) {
	data := SignABI(RedPacket_changeWallet) + packNumeric(address, 32)

	return hex.DecodeString(data)
}

func ChangeMaxCount(value string) ([]byte, error) {
	data := SignABI(RedPacket_changeMaxCount) + packNumeric(value, 32)

	return hex.DecodeString(data)
}

func GetRedPacketDetail(value string) ([]byte, error) {
	data := SignABI(RedPacket_getRedPacketDetail) + packNumeric(value, 32)

	return hex.DecodeString(data)
}

func NewRedPacket(address, from string, value, count, cmd string) ([]byte, error) {

	data := SignABI(RedPacket_newRedPacket) +
		packNumeric(address, 32) +
		packNumeric(from, 32) +
		packNumeric(value, 32) +
		packNumeric(count, 32) +
		packNumeric(cmd, 32)

	return hex.DecodeString(data)
}

func Open(tokeId, address string, cmd string) ([]byte, error) {

	data := SignABI(RedPacket_open) +
		packNumeric(tokeId, 32) +
		packNumeric(address, 32) +
		packNumeric(cmd, 32)

	return hex.DecodeString(data)
}

func OpenMany(tokeId string, addresses []string, cmd string) ([]byte, error) {

	start := hex.EncodeToString(big.NewInt(96).Bytes())

	data := SignABI(RedPacket_openMany) +
		packNumeric(tokeId, 32) +
		packNumeric(start, 32) +
		packNumeric(cmd, 32) +
		encodeStrings(addresses)

	return hex.DecodeString(data)
}

func SendEther(value string) ([]byte, error) {
	value = packNumeric(value, 32)

	data := SignABI(RedPacket_sendEther) + value

	return hex.DecodeString(data)
}

func TakeBack(tokeId string) ([]byte, error) {
	data := SignABI(RedPacket_takeBack) + packNumeric(tokeId, 32)

	return hex.DecodeString(data)
}

func encodeStrings(params []string) string {
	length := big.NewInt(int64(len(params)))

	lenStr := hex.EncodeToString(length.Bytes())

	codes := packNumeric(lenStr, 64)

	for _, v := range params {
		codes += packNumeric(v, 32)
	}

	return codes
}
