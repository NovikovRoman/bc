package bc

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_NewNews(t *testing.T) {
	n, err := NewNews(testdata)
	require.Nil(t, err)
	require.Len(t, n, 10)
	require.Equal(t, n[6].Title, "Добавлена новая функция")
	require.Equal(t, n[6].Date, "22 мая 2024, 19:42")
}
