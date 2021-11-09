package dht

import (
	"container/heap"
	"errors"
)

var (
	ErrBucketFull = errors.New("bucket is full")
)

const (
	BucketSize = 20
)

type Node struct {
	address Address
	buckets [256]bucket
}

// NewNode creates a new node.
func NewNode(address Address) (*Node, error) {
	return &Node{
		address: address,
		buckets: [256]bucket{},
	}, nil
}

// String returns the string representation of the node.
func (n *Node) String() string {
	return n.address.String()
}

// AddAddress adds the address to one of the Node's buckets.
func (n *Node) AddAddress(address Address) {
	bi := newUint256(address.bytes).leadingZeros()
	dist := Dist(n.address, address)
	n.buckets[bi].add(bucketEntry{
		addr: address,
		dist: dist,
	})
}

// Neighbors returns the k closest nodes to the Node.
func (n *Node) Neighbors(address Address) []Address {
	bi := newUint256(address.bytes).leadingZeros()
	addresses := make([]Address, 0, len(n.buckets[bi]))
	for _, be := range n.buckets[bi] {
		addresses = append(addresses, be.addr)
	}
	return addresses
}

// bucketEntry represents an entry in a bucket.
type bucketEntry struct {
	addr Address
	dist Distance
}

// bucket is a struct that handles the distance-sorted storage of addresses.
type bucket []bucketEntry

func (b bucket) Len() int {
	return len(b)
}

func (b bucket) Less(i, j int) bool {
	return b[j].dist.Less(b[i].dist)
}

func (b bucket) Swap(i, j int) {
	b[i], b[j] = b[j], b[i]
}

func (b *bucket) Push(x interface{}) {
	*b = append(*b, x.(bucketEntry))
}

func (b *bucket) Pop() interface{} {
	v := (*b)[len(*b)-1]
	*b = (*b)[:len(*b)-1]
	return v
}

func (b *bucket) add(be bucketEntry) {
	if len(*b) < BucketSize {
		heap.Push(b, be)
		return
	}

	if (*b)[0].dist.Less(be.dist) {
		return
	}

	(*b)[0] = be
	heap.Fix(b, 0)
}
