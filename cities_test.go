package bc

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_CitiesFromFile(t *testing.T) {
	c, err := NewCities(testdata)
	require.Nil(t, err)
	require.Len(t, c, 446)
}
