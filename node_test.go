package dht_test

import (
	"net"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/watnet/dht"
	"github.com/watnet/ecdh"
)

func TestNode(t *testing.T) {
	for i := 0; i < 10; i++ {
		var c ecdh.Config
		priv1, err := c.GenerateKey()
		require.NoError(t, err)
		pub1, err := priv1.PublicKey()
		require.NoError(t, err)
		a1 := dht.NewAddressV1(pub1)

		priv2, err := c.GenerateKey()
		require.NoError(t, err)
		pub2, err := priv2.PublicKey()
		a2 := dht.NewAddressV1(pub2)
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
