package fixed

import (
	"encoding/hex"
	"math"
	"math/big"
	"strings"
)

// Number fixed number
type Number struct {
	value    int64
	decimals int
}

func hexBytes(value string) ([]byte, error) {
	value = strings.TrimPrefix(value, "0x")

	if len(value)%2 != 0 {
		value = "0" + value
	}

	return hex.DecodeString(value)
}

// New .
func New(value int64, decimals int) *Number {
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
func (number *Number) Decimals() int {
	return number.decimals
}

// FromFloat .
func FromFloat(value *big.Float, decimals int) *Number {
	var val2 = big.NewInt(1)

	for i := uint64(0); i < uint64(math.Abs(float64(decimals))); i++ {
		val2 = new(big.Int).Mul(val2, big.NewInt(10))
	}

	number := new(Number)

	number.decimals = decimals

	if decimals > 0 {
		val := new(big.Float).Mul(value, new(big.Float).SetInt(val2))

		number.value, _ = val.Int64()
	} else {
		val := new(big.Float).Quo(value, new(big.Float).SetInt(val2))

		number.value, _ = val.Int64()
	}

	return number
}

// FromHex decode number from hex string
func FromHex(value string, decimals int) (*Number, error) {
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

	var val2 = big.NewInt(1)

	for i := uint64(0); i < uint64(math.Abs(float64(number.decimals))); i++ {
		val2 = new(big.Int).Mul(val2, big.NewInt(10))
	}

	if number.decimals > 0 {
		return new(big.Float).Quo(new(big.Float).SetInt(bigValue), new(big.Float).SetInt(val2))
	}

	return new(big.Float).Mul(new(big.Float).SetInt(bigValue), new(big.Float).SetInt(val2))
}

// Compare return > 0  if self > to return < 0  if self < to  return = 0 if self == to
func (number *Number) Compare(to *Number) int {
	return new(big.Float).Sub(number.Float(), to.Float()).Sign()
}
