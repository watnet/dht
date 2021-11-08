package dht

import (
	"crypto/rand"
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

// NewAddress creates a new address random address.
// The bytes will be generated according to the supplied version.
func NewAddress(version int) (Address, error) {
	switch version {
	case 1:
		var result [32]byte
		_, err := rand.Read(result[:32])
		if err != nil {
			panic(err)
		}
		return Address{version: version, bytes: result[:]}, nil
	default:
		return Address{}, fmt.Errorf("unknown version %v", version)
	}
}

// NewAddressV1 creates a new version 1 address from the given bytes.
func NewAddressV1(bytes []byte) Address {
	return Address{version: 1, bytes: bytes}
}

func (a Address) Version() int {
	return a.version
}

// String returns the string representation of the address in the following format:
//    wn<version>.<encoded bytes>
// For example:
//    wn1.ybndrfg8ejkmcpqxot1uwisza345h769
func (a Address) String() string {
	enc := zEncoding.EncodeToString(a.bytes)
	return fmt.Sprintf("wn%v%c%v", a.version, VersionSeparator, enc)
}

// ParseAddress parses a string representation of an address into an address.
func ParseAddress(s string) (a Address, err error) {
	// TODO: Should't this parse differently according to the version?
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
