package dht_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/watnet/dht"
)

func TestAddress(t *testing.T) {
	// FIXME: This test needs to be totally reworked in order to work with the updated ecdh package.
	randAddr, err := dht.NewAddressV1()
	require.NoError(t, err)
	require.Equal(t, 1, randAddr.Version())
	require.Equal(t, 56, len(randAddr.String()))

	const addr1 = "wn1.ktwg1h3ypf31yamqrbose3d1ci3zgedf"
	a1, err := dht.ParseAddress(addr1)
	require.NoError(t, err, "It should return no error.")
	require.Equal(t, addr1, a1.String())

	const addr2 = "wn1.ktwg1h3ypf31yajyctwsc3ufqj1sh7df"
	a2, err := dht.ParseAddress(addr2)
	require.NoError(t, err, "It should return no error.")
	require.Equal(t, addr2, a2.String())
}

// func TestExample(t *testing.T) {
// 	tests := []struct {
// 		name string
// 		in   string
// 		out  int
// 	}{
// 		{
// 			name: "addr1",
// 			in:   "wn1.ktwg1h3ypf31yamqrbose3d1ci3zgedf",
// 			out:  1,
// 		},
// 		{
// 			name: "addr2",
// 			in:   "wn1.ktwg1h3ypf31yajyctwsc3ufqj1sh7df",
// 			out:  1,
// 		},
// 	}

// 	for _, test := range tests {
// 		t.Run(test.name, func(t *testing.T) {
// 			got, err := dht.ParseAddress(test.in)
// 			if err != nil {
// 				t.Errorf("ParseAddress(%q) = %v", test.in, err)
// 			} else {
// 				require.NoError(t, err, "It should return no error.")
// 			}
// 			require.Equal(t, test.out, got.String(), "The expected output should equal the actual output.")
// 			require.Equal(t, test.out, got.Version(), "The expected output should equal the actual output.")
// 		})
// 	}
// }
