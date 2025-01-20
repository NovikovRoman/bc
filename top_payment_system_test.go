package bc

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNewTopPaymentSystems(t *testing.T) {
	var tps []TopPaymentSystem
	tps, err := NewTopPaymentSystems(testdata)
	require.Nil(t, err)
	require.Len(t, tps, 100)
}
