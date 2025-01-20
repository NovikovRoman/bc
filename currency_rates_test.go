package bc

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_NewCurrencyRates(t *testing.T) {
	r, err := NewCurrencyRates(testdata)
	require.Nil(t, err)
	require.Len(t, r, 104)
	require.Len(t, r[840], 104)
}
