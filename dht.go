package dht

type Dht struct {
}

type Address struct {
	bytes []byte
}

func Distance(a1, a2 Address) []byte {
	var result [32]byte

	for i := range a1.bytes {
		result[i] = a1.bytes[i] ^ a2.bytes[i]
	}

	return result[:]
}
