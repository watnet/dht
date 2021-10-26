package dht_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/watnet/dht"
)

func TestNode(t *testing.T) {
	const addr1 = "wn1.ktwg1h3ypf31yamqrbose3d1ci3zgedf"
	a1, _ := dht.ParseAddress(addr1)
	const addr2 = "wn1.ktwg1h3ypf31yajyctwsc3ufqj1sh7df"
	a2, _ := dht.ParseAddress(addr2)

	var a3 = dht.NewAddress(1, dht.Distance(a1, a2))
	// a3.String() -> wn1.yyyyyyyyyyyyyynqeoryryoznhmb4iyyyyyyyyyyyyyyyyyyyyyy
	// []byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 78, 68, 8, 2, 2, 23, 23, 22, 29, 84, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
	// zeroes: 73

	n, _ := dht.NewNode(a1)
	require.Equal(t, 0, n.Size())

	err := n.AddAddress(a3)
	require.NoError(t, err)

	require.Equal(t, 1, n.Size())
}
