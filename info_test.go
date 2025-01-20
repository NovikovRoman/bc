package bc

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_NewInfo(t *testing.T) {
	i, err := NewInfo(testdata)
	require.Nil(t, err)
	require.Equal(t, i.LastUpdate.Format("2006-01-02 15:04:05"), "2025-01-17 15:07:37")
	require.Equal(t, i.CompatibleVersion, "2.02")
	require.Equal(t, i.CurrentVersion, "2.03")
}
