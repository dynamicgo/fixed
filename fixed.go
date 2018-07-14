package fixed

import (
	"encoding/hex"
	"math/big"
	"strings"
)

// Number fixed number
type Number struct {
	value    int64
	decimals uint
}

func hexBytes(value string) ([]byte, error) {
	value = strings.TrimPrefix(value, "0x")

	if len(value)%2 != 0 {
		value = "0" + value
	}

	return hex.DecodeString(value)
}

// New .
func New(value int64, decimals uint) *Number {
	return &Number{
		value:    value,
		decimals: decimals,
	}
}

// Value .
func (number *Number) Value() int64 {
	return number.value
}

// HexValue return value as hex string
func (number *Number) HexValue() string {
	return "0x" + hex.EncodeToString(big.NewInt(number.value).Bytes())
}

// Decimals .
func (number *Number) Decimals() uint {
	return number.decimals
}

// FromHex decode number from hex string
func FromHex(value string, decimals uint) (*Number, error) {
	valueBytes, err := hexBytes(value)

	if err != nil {
		return nil, err
	}

	bigValue := new(big.Int).SetBytes(valueBytes)

	return &Number{
		value:    bigValue.Int64(),
		decimals: decimals,
	}, nil
}

// Float convert to big.Float
func (number *Number) Float() *big.Float {
	bigValue := big.NewInt(number.value)
	bigDecimals := big.NewInt(int64(number.decimals))

	var val2 = big.NewInt(1)

	for i := uint64(0); i < bigDecimals.Uint64(); i++ {
		val2 = new(big.Int).Mul(val2, big.NewInt(10))
	}

	return new(big.Float).Quo(new(big.Float).SetInt(bigValue), new(big.Float).SetInt(val2))
}

// Compare return > 0  if self > to return < 0  if self < to  return = 0 if self == to
func (number *Number) Compare(to *Number) int {
	return new(big.Float).Sub(number.Float(), to.Float()).Sign()
}
