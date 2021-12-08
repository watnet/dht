package dht_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/watnet/dht"
	"github.com/watnet/ecdh"
)

func TestAddress(t *testing.T) {
	for i := 0; i < 10; i++ {
		var c ecdh.Config
		priv, err := c.GenerateKey()
		require.NoError(t, err)
		pub, err := priv.PublicKey()
		require.NoError(t, err)
		addr := dht.NewAddressV1(pub)
		require.Equal(t, 1, addr.Version())
		require.Equal(t, 56, len(addr.String()))

		a, err := dht.ParseAddress(addr.String())
		require.NoError(t, err)
		require.Equal(t, addr.String(), a.String())
	}
}
