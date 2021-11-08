package dht_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/watnet/dht"
)

func TestNode(t *testing.T) {
	for i := 0; i < 10; i++ {
		fmt.Printf("%d\n", i)
		a1, err := dht.NewAddress(1)
		require.NoError(t, err)
		a2, err := dht.NewAddress(1)
		require.NoError(t, err)

		n1, err := dht.NewNode(a1, 20)
		require.NoError(t, err)
		n2, err := dht.NewNode(a2, 20)
		require.NoError(t, err)

		// FIXME: This test does not work. It seems to be breaking because the Node's bucket
		// capacity is 0 instead of 20 (sometimes), so it instantly fails with a bucket full
		// error. I'm assuming that the bucket is full because of how the bucket is
		// implemented inside the Node:
		//     buckets: [256]bucket{
		//         {
		// 		   addresses: make([]Address, 0, bucketCapacity),
		// 	       },
		//     },

		err = n1.AddAddress(a2)
		require.NoError(t, err)
		err = n2.AddAddress(a1)
		require.NoError(t, err)
	}
}

func TestFullBucket(t *testing.T) {
	a1, err := dht.NewAddress(1)
	require.NoError(t, err)
	a2, err := dht.NewAddress(1)
	require.NoError(t, err)

	n, err := dht.NewNode(a1, 20)
	require.NoError(t, err)

	for i := 0; i < 20; i++ {
		err := n.AddAddress(a2)
		require.NoError(t, err)
	}

	err = n.AddAddress(a2)
	require.Error(t, err)
}
