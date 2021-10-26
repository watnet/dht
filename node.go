package dht

import (
	"fmt"
	"math/bits"
)

type Node struct {
	address Address
	buckets bucketManager
}

// NewNode creates a new node.
func NewNode(address Address) (*Node, error) {
	return &Node{
		address: address,
		buckets: bucketManager{
			maxBucketSize: 20,
		},
	}, nil
}

// String returns the string representation of the node.
func (n *Node) String() string {
	return n.address.String()
}

// AddAddress adds the address to the Node's bucketManager.
// It is intended to be used with the result of  dht.Distance. For example:
//    node.AddAddress(dht.Distance(addr1, addr2))
func (n *Node) AddAddress(address Address) error {
	n.buckets.add(address)
	return nil
}

// Size returns the number of addresses in the Node's buckets.
func (n *Node) Size() int {
	// TODO: Node.Size is only here for testing. This way there's some way to check if the AddAddress
	// method is working correctly. Same with bucketManager.size down below.
	return n.buckets.size()
}

// bucketManager is a struct that manages the buckets of a node.
type bucketManager struct {
	// TODO: Currently, maxBucketSize is hard-coded to 20 above inside NewNode. There's *probably* a better way to do this...
	maxBucketSize int
	buckets       [256][]Address
}

// size returns the number of addresses in the bucketManager.
func (bm *bucketManager) size() int {
	size := 0
	for _, bucket := range bm.buckets {
		size += len(bucket)
	}
	return size
}

// add adds the address to the bucketManager's buckets.
func (bm *bucketManager) add(address Address) {
	var index int
	for i := range address.bytes {
		if address.bytes[i] != 0 {
			index = i
			break
		}
	}

	bi := 8*index + bits.LeadingZeros8(address.bytes[index])
	if len(bm.buckets[bi]) >= bm.maxBucketSize {
		bm.replaceInBucket(bm.buckets[bi], address)
		return
	}
	bm.buckets[bi] = append(bm.buckets[bi], address)
	fmt.Println("LOG: ", bm.buckets[bi])
}

// replaceInBucket replaces the lowest-priority address in the bucket with a new address.
// If no existing address is lower priority, this function is a no-op.
func (bm *bucketManager) replaceInBucket(bucket []Address, address Address) {
	for i, v := range bucket {
		if v.Less(address) {
			bucket[i] = address
			return
		}
	}
}
