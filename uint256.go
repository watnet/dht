package dht

import (
	"encoding/binary"
	"math/bits"
)

// uint256 represents a 256 bit unsigned integer.
type uint256 struct {
	hi  uint64
	mhi uint64
	mlo uint64
	lo  uint64
}

// isZero returns true if u == 0
func (u uint256) isZero() bool {
	return u.hi|u.mhi|u.mlo|u.lo == 0
}

// and returns the bitwise and of u and x (u&x)
func (u uint256) and(x uint256) uint256 {
	return uint256{
		u.hi & x.hi,
		u.mhi & x.mhi,
		u.mlo & x.mlo,
		u.lo & x.lo,
	}
}

// or returns the bitwise or of u and x (u|x)
func (u uint256) or(x uint256) uint256 {
	return uint256{
		u.hi | x.hi,
		u.mhi | x.mhi,
		u.mlo | x.mlo,
		u.lo | x.lo,
	}
}

// xor returns the bitwise xor of u and x (u^x)
func (u uint256) xor(x uint256) uint256 {
	return uint256{
		u.hi ^ x.hi,
		u.mhi ^ x.mhi,
		u.mlo ^ x.mlo,
		u.lo ^ x.lo,
	}
}

// not returns the bitwise not of u (^u)
func (u uint256) not() uint256 {
	return uint256{
		^u.hi,
		^u.mhi,
		^u.mlo,
		^u.lo,
	}
}

// add returns u+x
func (u uint256) add(x uint256) uint256 {
	lo, loCarry := bits.Add64(u.lo, x.lo, 0)
	mlo, mloCarry := bits.Add64(u.mlo, x.mlo, loCarry)
	mhi, mhiCarry := bits.Add64(u.mhi, x.mhi, mloCarry)
	return uint256{
		u.hi + x.hi + mhiCarry,
		mhi,
		mlo,
		lo,
	}
}

// sub returns u-x
func (u uint256) sub(x uint256) uint256 {
	lo, loBorrow := bits.Sub64(u.lo, x.lo, 0)
	mlo, mloBorrow := bits.Sub64(u.mlo, x.mlo, loBorrow)
	mhi, mhiBorrow := bits.Sub64(u.mhi, x.mhi, mloBorrow)
	return uint256{
		u.hi - x.hi - mhiBorrow,
		mhi,
		mlo,
		lo,
	}
}

// subOne returns u-1
func (u uint256) subOne() uint256 {
	return u.sub(uint256{0, 0, 0, 1})
}

// addOne returns u+1
func (u uint256) addOne() uint256 {
	return u.add(uint256{0, 0, 0, 1})
}

// leadingZeros returns the number of leading zero bits in u.
func (u uint256) leadingZeros() int {
	if u.hi != 0 {
		return bits.LeadingZeros64(u.hi)
	}
	if u.mhi != 0 {
		return bits.LeadingZeros64(u.mhi) + 64
	}
	if u.mlo != 0 {
		return bits.LeadingZeros64(u.mlo) + 128
	}
	return bits.LeadingZeros64(u.lo) + 192
}

// less returns true if u < x.
func (u uint256) less(x uint256) bool {
	if u.hi < x.hi {
		return true
	}
	if u.mhi < x.mhi {
		return u.hi == x.hi
	}
	if u.mlo < x.mlo {
		return u.mhi == x.mhi && u.hi == x.hi
	}
	return u.lo < x.lo && u.mlo == x.mlo && u.mhi == x.mhi && u.hi == x.hi
}

// newUint256 converts a byte slice to a uint256.
func newUint256(b []byte) uint256 {
	return uint256{
		binary.BigEndian.Uint64(b[:8]),
		binary.BigEndian.Uint64(b[8:16]),
		binary.BigEndian.Uint64(b[16:24]),
		binary.BigEndian.Uint64(b[24:]),
	}
}

// Bytes converts a uint256 to a byte slice.
func (u uint256) Bytes() []byte {
	b := make([]byte, 32)
	binary.BigEndian.PutUint64(b[:8], u.hi)
	binary.BigEndian.PutUint64(b[8:16], u.mhi)
	binary.BigEndian.PutUint64(b[16:24], u.mlo)
	binary.BigEndian.PutUint64(b[24:], u.lo)
	return b
}
