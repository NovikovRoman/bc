package bc

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNewExchangeRates(t *testing.T) {
	e, err := NewExchangeRates(testdata)
	require.Nil(t, err)
	require.Len(t, e, 486436)
}
