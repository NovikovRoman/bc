package bc

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_CurrenciesFromFile(t *testing.T) {
	c, err := NewCurrencies(testdata)
	require.Nil(t, err)
	require.Len(t, c, 104)
}
