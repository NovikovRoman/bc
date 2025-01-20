package bc

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_NewPaymentSystems(t *testing.T) {
	ps, err := NewPaymentSystems(testdata)
	require.Nil(t, err)
	require.Len(t, ps, 322)
	require.Equal(t, ps[0].ID, 93)
	require.Equal(t, ps[0].Name, "Bitcoin (BTC)")
	require.Equal(t, ps[0].NameAlt, "BTC")
	require.Equal(t, ps[0].Type, CryptoType)
	require.Equal(t, ps[0].CurrencyID, 1050)
}
