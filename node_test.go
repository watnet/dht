package dht_test

import (
	"net"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/watnet/dht"
)

func TestNode(t *testing.T) {
	for i := 0; i < 10; i++ {
		a1, err := dht.NewAddress(1)
		require.NoError(t, err)
		a2, err := dht.NewAddress(1)
		require.NoError(t, err)

		n1, err := dht.NewNode(a1)
		require.NoError(t, err)
		n2, err := dht.NewNode(a2)
		require.NoError(t, err)

		// randIPs is a slice of random net.IPAddr
		randIPs1 := make([]net.IPAddr, 0, 20)
		randIPs2 := make([]net.IPAddr, 0, 20)

		n1.Add(a2, randIPs1)
		n2.Add(a1, randIPs2)
	}
}
