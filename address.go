package dht

import (
	"encoding/base32"
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

func (a Address) String() string {
	enc := zEncoding.EncodeToString(a.bytes)
	return fmt.Sprintf("wn%v%c%v", a.version, VersionSeparator, enc)
}

func (a Address) MarshalText() (data []byte, err error) {
	return []byte(a.String()), nil
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
