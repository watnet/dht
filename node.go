package dht

import (
	"container/heap"
	"net"
	"time"
)

const (
	BucketSize  = 20
	AddrTimeout = 24 * time.Hour
)

type Node struct {
	address Address
	buckets [256]bucket
}

// NewNode creates a new node.
func NewNode(address Address) (*Node, error) {
	return &Node{
		address: address,
	}, nil
}

// String returns the string representation of the node.
func (n *Node) String() string {
	return n.address.String()
}

// Add adds the address to one of the Node's buckets.
func (n *Node) Add(addr Address, ip []net.IPAddr) {
	bi := newUint256(addr.bytes).leadingZeros()
	dist := Dist(n.address, addr)
	n.buckets[bi].add(addr, dist, ip)
}

// Neighbors returns the k closest nodes to the Node.
func (n *Node) Neighbors(address Address) []PeerInfo {
	bi := newUint256(address.bytes).leadingZeros()
	addresses := make([]PeerInfo, 0, len(n.buckets[bi]))
	for _, be := range n.buckets[bi] {
		addresses = append(addresses, be.addr)
	}
	return addresses
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

func (b *bucket) add(address Address, dist Distance, ip []net.IPAddr) {
	be := bucketEntry{
		addr:  PeerInfo{address, ip},
		dist:  dist,
		added: time.Now(),
	}

	if len(*b) < BucketSize {
		heap.Push(b, be)
		return
	}

	oldest := 0
	for i := 1; i < len(*b); i++ {
		if (*b)[i].added.Before((*b)[oldest].added) {
			oldest = i
		}
	}
	if (*b)[oldest].age() > AddrTimeout {
		(*b)[oldest] = be
		heap.Fix(b, oldest)
		return
	}

	if (*b)[0].dist.Less(be.dist) {
		return
	}

	(*b)[0] = be
	heap.Fix(b, 0)
}

// bucketEntry represents an entry in a bucket.
type bucketEntry struct {
	addr  PeerInfo
	dist  Distance
	added time.Time
}

// age returns the age of the entry.
func (be bucketEntry) age() time.Duration {
	return time.Since(be.added)
}
