package dht

import (
	"errors"
	"fmt"
)

var (
	ErrBucketFull = errors.New("bucket is full")
)

type Node struct {
	address Address
	buckets [256]bucket
}

// NewNode creates a new node.
func NewNode(address Address, bucketCapacity int) (*Node, error) {
	return &Node{
		address: address,
		buckets: [256]bucket{
			{
				addresses: make([]Address, 0, bucketCapacity),
			},
		},
	}, nil
}

// String returns the string representation of the node.
func (n *Node) String() string {
	return n.address.String()
}

// AddAddress adds the address to one of the Node's buckets.
func (n *Node) AddAddress(address Address) error {
	bi := newUint256(address.bytes).leadingZeros()
	err := n.buckets[bi].add(n.address, address)
	if err != nil {
		return fmt.Errorf("cannot add address: %v", err)
	}
	return nil
}

// Closest returns the closest address to the target address.
func (n *Node) Closest(address Address) Address {
	bi := newUint256(address.bytes).leadingZeros()
	for _, v := range n.buckets[bi].addresses {
		if Dist(n.address, v).Less(Dist(n.address, address)) {
			return v
		}
	}
	return n.buckets[bi].addresses[0]
}

// bucket is a struct that handles the distance-sorted storage of addresses.
type bucket struct {
	addresses []Address
}

// add inserts the address into the proper index of the bucket, based on the distance to the target address.
func (b *bucket) add(nodeaddress, address Address) error {
	for i, v := range b.addresses {
		if Dist(nodeaddress, v).Less(Dist(nodeaddress, address)) {
			b.addresses[i] = address
			return nil
		}
	}

	if len(b.addresses) < cap(b.addresses) {
		b.addresses = append(b.addresses, address)
		return nil
	}

	return ErrBucketFull
}
