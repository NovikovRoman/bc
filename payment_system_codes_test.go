package bc

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_NewPaymentSystemCodes(t *testing.T) {
	ps, err := NewPaymentSystemCodes(testdata)
	require.Nil(t, err)
	require.Len(t, ps, 322)
	require.Equal(t, ps[93], "BTC")
}
