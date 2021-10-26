package dht_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/watnet/dht"
)

func TestAddress(t *testing.T) {
	const addr1 = "wn1.ktwg1h3ypf31yamqrbose3d1ci3zgedf"
	a1, err := dht.ParseAddress(addr1)
	require.NoError(t, err, "It should return no error.")
	require.Equal(t, addr1, a1.String())

	const addr2 = "wn1.ktwg1h3ypf31yajyctwsc3ufqj1sh7df"
	a2, err := dht.ParseAddress(addr2)
	require.NoError(t, err, "It should return no error.")
	require.Equal(t, addr2, a2.String())
}
