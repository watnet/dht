package dht

// Distance represents a distance between two addresses.
type Distance struct {
	v uint256
}

// Less returns true if the distance is less than the other Distance.
func (d Distance) Less(other Distance) bool {
	return d.v.less(other.v)
}

// Dist returns the distance between two Addresses as a new Distance.
//
// The order of the arguments has no effect on the output, so that:
//    Dist(a1, a2) == Dist(a2, a1)
func Dist(a1, a2 Address) Distance {
	return Distance{newUint256(a1.bytes).xor(newUint256(a2.bytes))}
}
