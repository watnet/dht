package dht

import (
	"encoding/base32"
	"encoding/binary"
	"errors"
	"fmt"
	"strconv"
	"strings"
)

const (
	VersionSeparator = '.'
)

var (
	ErrNoSeparator = errors.New("no separator")
)

// zEncoding is a z-base-32 encoding scheme.
var zEncoding = base32.NewEncoding("ybndrfg8ejkmcpqxot1uwisza345h769").WithPadding(base32.NoPadding)

type Address struct {
	version int
	bytes   []byte
}

// NewAddress creates a new address.
func NewAddress(version int, bytes []byte) Address {
	// TODO: Is bytes going to be randomely generated?
	return Address{version: version, bytes: bytes}
}

func ParseAddress(s string) (a Address, err error) {
	sep := strings.LastIndex(s, string(VersionSeparator))
	if sep < 0 {
		return a, ErrNoSeparator
	}
	if len(s) <= sep {
		return a, fmt.Errorf("length %v less than sep index %v", len(s), sep)
	}

	v, err := strconv.ParseInt(s[2:sep], 10, 0)
	if err != nil {
		return a, fmt.Errorf("parse version: %w", err)
	}

	dec, err := zEncoding.DecodeString(s[sep+1:])
	if err != nil {
		return a, fmt.Errorf("decode address: %w", err)
	}

	return Address{
		version: int(v),
		bytes:   dec,
	}, nil
}

// Version returns the address's version field.
func (a Address) Version() int {
	return a.version
}

// String returns the string representation of the address.
func (a Address) String() string {
	enc := zEncoding.EncodeToString(a.bytes)
	return fmt.Sprintf("wn%v%c%v", a.version, VersionSeparator, enc)
}

// MarshalText implements encoding.TextMarshaler.
func (a Address) MarshalText() (data []byte, err error) {
	return []byte(a.String()), nil
}

// UnmarshalText implements encoding.TextUnmarshaler.
func (a *Address) UnmarshalText(data []byte) error {
	addr, err := ParseAddress(string(data))
	if err != nil {
		return err
	}
	*a = addr
	return nil
}

// MarshalBinary implements encoding.BinaryMarshaler.
func (a Address) MarshalBinary() (data []byte, err error) {
	data = make([]byte, 8+len(a.bytes))
	vlen := binary.PutUvarint(data, uint64(a.version))
	copy(data[vlen:], a.bytes)
	return data[:vlen+len(a.bytes)], nil
}

// UnmarshalBinary implements encoding.BinaryUnmarshaler.
func (a *Address) UnmarshalBinary(data []byte) error {
	if len(data) < 8 {
		return fmt.Errorf("length %v less than 8", len(data))
	}
	v, n := binary.Uvarint(data)
	if n <= 0 {
		return fmt.Errorf("parse version: %w", ErrNoSeparator)
	}
	a.version = int(v)
	a.bytes = data[n:]
	return nil
}

// Less returns true if the Address is less than the other address.
func (a Address) Less(b Address) bool {
	// TODO: Review this method and see if it really does what it should.
	return a.String() < b.String()
}

// Distance calculates the distance between two Addresses.
//
// The order of the arguments has no effect on the output, so that:
//    Distance(a1, a2) == Distance(a2, a1)
func Distance(a1, a2 Address) []byte {
	var result [32]byte

	for i := range a1.bytes {
		result[i] = a1.bytes[i] ^ a2.bytes[i]
	}

	return result[:]
}
