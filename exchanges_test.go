package bc

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_NewExchanges(t *testing.T) {
	e, err := NewExchanges(testdata)
	require.Nil(t, err)
	require.Len(t, e, 527)
}
