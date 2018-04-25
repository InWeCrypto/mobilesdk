package erc721

import (
	"encoding/hex"
	"fmt"
	"math/big"
	"strings"

	"github.com/inwecrypto/sha3"
)

const (
	ownerOf             = "ownerOf(uint256)"
	setApprovalForAll   = "setApprovalForAll(address,bool)"
	getApprovedAddress  = "getApprovedAddress(uint256)"
	isApprovedForAll    = "isApprovedForAll(address,address)"
	takeOwnership       = "takeOwnership(uint256)"
	tokenOfOwnerByIndex = "tokenOfOwnerByIndex(address,uint256)"
	tokenMetadata       = "tokenMetadata(uint256)"
	tokensOf            = "tokensOf(address)"
	exists              = "exists(uint256)"
	setAssetHolder      = "setAssetHolder(address,uint256)"
	isAuthorized        = "isAuthorized(address,uint256)"
	description         = "description()"

	// DecentraLand
	DecentraLand_decodeTokenId = "decodeTokenId(uint256)"
	DecentraLand_encodeTokenId = "encodeTokenId(int256,int256)"
	DecentraLand_landData      = "landData(int256,int256)"
	DecentraLand_landOf        = "landOf(address)"
	DecentraLand_transferLand  = "transferLand(int,int,address)"
	DecentraLand_ownerOfLand   = "ownerOfLand(int,int)"

	// RedPacket
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
	RedPacket_maxCount           = "maxCount()"
)

// Method/Event id
var (
	Method_ownerOf             = SignABI(ownerOf)
	Method_setApprovalForAll   = SignABI(setApprovalForAll)
	Method_getApprovedAddress  = SignABI(getApprovedAddress)
	Method_isApprovedForAll    = SignABI(isApprovedForAll)
	Method_takeOwnership       = SignABI(takeOwnership)
	Method_tokenOfOwnerByIndex = SignABI(tokenOfOwnerByIndex)
	Method_tokenMetadata       = SignABI(tokenMetadata)
	Method_tokensOf            = SignABI(tokensOf)
	Method_exists              = SignABI(exists)
	Method_setAssetHolder      = SignABI(setAssetHolder)
	Method_isAuthorized        = SignABI(isAuthorized)
	Method_description         = SignABI(description)
)

// SignABI sign abi string
func SignABI(abi string) string {
	hasher := sha3.NewKeccak256()
	hasher.Write([]byte(abi))
	data := hasher.Sum(nil)

	return hex.EncodeToString(data[0:4])
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

func packNumeric(value string, bytes int) string {
	value = strings.TrimPrefix(value, "0x")

	chars := bytes * 2

	n := len(value)
	if n%chars == 0 {
		return value
	}
	return strings.Repeat("0", chars-n%chars) + value
}

func TransferLand(to string, x, y string) ([]byte, error) {
	to = packNumeric(to, 32)
	x = packNumeric(x, 32)
	y = packNumeric(y, 32)

	data := fmt.Sprintf("%s%s%s%s", SignABI(DecentraLand_transferLand), x, y, to)

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

	return fmt.Sprintf("0x%s%s", SignABI(DecentraLand_decodeTokenId), value)
}

func EncodeTokenId(x, y string) string {
	x = packNumeric(x, 32)
	y = packNumeric(y, 32)

	return fmt.Sprintf("0x%s%s%s", SignABI(DecentraLand_encodeTokenId), x, y)
}

func LandData(x, y string) string {
	x = packNumeric(x, 32)
	y = packNumeric(y, 32)

	return fmt.Sprintf("0x%s%s%s", SignABI(DecentraLand_landData), x, y)
}

func Description() string {
	return fmt.Sprintf("0x%s", Method_description)
}

func LandOf(address string) string {
	address = packNumeric(address, 32)

	return fmt.Sprintf("0x%s%s", SignABI(DecentraLand_landOf), address)
}

func OwnerOfLand(x, y string) string {
	x = packNumeric(x, 32)
	y = packNumeric(y, 32)

	return fmt.Sprintf("0x%s%s%s", SignABI(DecentraLand_ownerOfLand), x, y)
}

func TaxCost() string {
	return "0x" + SignABI(RedPacket_taxCost)
}

func MaxCount() string {
	return "0x" + SignABI(RedPacket_maxCount)
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

func GetRedPacketDetail(value string) string {
	data := "0x" + SignABI(RedPacket_getRedPacketDetail) + packNumeric(value, 32)

	return data
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
