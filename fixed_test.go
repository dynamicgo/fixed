package fixed

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSign(t *testing.T) {

	number, err := FromHex("d42e876228795e400000", 18)

	println(number.Float().String())

	require.NoError(t, err)

	// require.True(t, number.Compare(FromFloat(big.NewFloat(0), 0)) < 0)

	// require.Equal(t, number.HexValue(), "-100000000")

	// number = FromFloat(big.NewFloat(-0.111), 2)

	// require.True(t, number.Compare(FromFloat(big.NewFloat(0), 0)) < 0)
}
