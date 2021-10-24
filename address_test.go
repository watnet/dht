package dht_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/watnet/dht"
)

func TestAddress(t *testing.T) {
	const addr = "wn1.ybndrfg8ejkmc"
	a, err := dht.ParseAddress(addr)
	require.NoError(t, err, "It should return no error.")
	require.Equal(t, addr, a.String())
}
