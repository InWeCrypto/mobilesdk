package ethmobile

import (
	"encoding/json"

	"github.com/inwecrypto/ethgo/erc20"
	"github.com/inwecrypto/ethgo/erc721"
	"github.com/inwecrypto/ethgo/rpc"
)

type EthCall struct {
}

func NewEthCall() *EthCall {
	return &EthCall{}
}

func (self *EthCall) Call(contract string, data string) (string, error) {
	site := rpc.CallSite{
		To:   contract,
		Data: data,
	}

	b, err := json.Marshal(site)
	if err != nil {
		return "", err
	}

	return string(b), nil
}

func (self *EthCall) Decimals(contract string) (string, error) {
	data := erc20.GetDecimals()

	return self.Call(contract, data)
}

func (self *EthCall) TotalSupply(contract string) (string, error) {
	data := erc20.GetTotalSupply()

	return self.Call(contract, data)
}

func (self *EthCall) BalanceOf(contract string, address string) (string, error) {
	data := erc20.BalanceOf(address)

	return self.Call(contract, data)
}

func (self *EthCall) Name(contract string) (string, error) {
	data := erc20.GetName()

	return self.Call(contract, data)
}

func (self *EthCall) LandDecodeTokenId(contract string, value string) (string, error) {
	data := erc721.DecodeTokenId(value)

	return self.Call(contract, data)
}

func (self *EthCall) LandEncodeTokenId(contract string, x, y string) (string, error) {
	data := erc721.EncodeTokenId(x, y)

	return self.Call(contract, data)
}

func (self *EthCall) LandData(contract string, x, y string) (string, error) {
	data := erc721.LandData(x, y)

	return self.Call(contract, data)
}

func (self *EthCall) LandOf(contract string, address string) (string, error) {
	data := erc721.LandOf(address)

	return self.Call(contract, data)
}

func (self *EthCall) OwnerOfLand(contract string, x, y string) (string, error) {
	data := erc721.OwnerOfLand(x, y)

	return self.Call(contract, data)
}

func (self *EthCall) Description(contract string) (string, error) {
	data := erc721.Description()

	return self.Call(contract, data)
}

func (self *EthCall) TokensOf(contract string, address string) (string, error) {
	data := erc721.TokensOf(address)

	return self.Call(contract, data)
}

func (self *EthCall) Exists(contract string, value string) (string, error) {
	data := erc721.IsExists(value)

	return self.Call(contract, data)
}

func (self *EthCall) TokenMetadata(contract string, value string) (string, error) {
	data := erc721.GetTokenMetadata(value)

	return self.Call(contract, data)
}

func (self *EthCall) TokenOfOwnerByIndex(contract string, address string, value string) (string, error) {
	data := erc721.TokenOfOwnerByIndex(address, value)

	return self.Call(contract, data)
}

func (self *EthCall) OwnerOf(contract string, value string) (string, error) {
	data := erc721.OwnerOf(value)

	return self.Call(contract, data)
}

func (self *EthCall) RedPacketTaxCost(contract string) (string, error) {
	data := erc721.TaxCost()

	return self.Call(contract, data)
}

func (self *EthCall) RedPacketMaxCount(contract string) (string, error) {
	data := erc721.MaxCount()

	return self.Call(contract, data)
}

func (self *EthCall) RedPacketDetail(contract string, value string) (string, error) {
	data := erc721.GetRedPacketDetail(value)

	return self.Call(contract, data)
}
