package dht

type address struct {
	version string
	bytes   []byte
}

// Version returns the address's version field.
func (a address) Version() string {
	return a.version
}

// Bytes returns the address's bytes field.
func (a address) Bytes() []byte {
	return a.bytes
}

// Distance calculates the XORed value of two addresses.
func Distance(a1, a2 address) []byte {
	var result [32]byte

	for i := range a1.Bytes() {
		result[i] = a1.Bytes()[i] ^ a2.Bytes()[i]
	}

	return result[:]
}
