package ethgo

import (
	"fmt"
	"math/big"
)

// Unit  .
type Unit float64

// These are the multipliers for ether denominations.
// Example: To get the wei value of an amount in 'douglas', use
//
//    new(big.Int).Mul(value, big.NewInt(params.Douglas))
//
const (
	Wei      Unit = 1
	Ada           = 1000 * Wei
	Babbage       = 1000 * Ada
	Shannon       = 1000 * Babbage
	Szabo         = 1000 * Shannon
	Finney        = 1000 * Szabo
	Ether         = 1000 * Finney
	Einstein      = 1000 * Ether
	Douglas       = 1000 * Einstein
)

// Value .
type Value big.Int

// As .
func (value *Value) As(unit Unit) *big.Float {
	return new(big.Float).Quo(new(big.Float).SetInt((*big.Int)(value)), big.NewFloat(float64(unit)))
}

// Format .
func (value *Value) Format(s fmt.State, format rune) {
	(*big.Int)(value).Format(s, format)
}

// Bytes .
func (value *Value) Bytes() []byte {
	return (*big.Int)(value).Bytes()
}

// NewValue create new eth value
func NewValue(val *big.Float, unit Unit) *Value {
	u := big.NewFloat(float64(unit))
	result, _ := new(big.Float).Mul(val, u).Int(nil)

	return (*Value)(result)
}

// CustomerValue .
func CustomerValue(val *big.Int, decimals *big.Int) *big.Float {

	var val2 = big.NewInt(10)

	for i := uint64(1); i < decimals.Uint64(); i++ {
		val2 = new(big.Int).Mul(val2, big.NewInt(10))
	}

	return new(big.Float).Quo(new(big.Float).SetInt(val), new(big.Float).SetInt(val2))
}

// FromCustomerValue .
func FromCustomerValue(val *big.Float, decimals *big.Int) *big.Int {
	var val2 = big.NewInt(10)

	for i := uint64(1); i < decimals.Uint64(); i++ {
		val2 = new(big.Int).Mul(val2, big.NewInt(10))
	}

	val3, _ := new(big.Float).Mul(val, new(big.Float).SetInt(val2)).Int(nil)

	return val3
}
