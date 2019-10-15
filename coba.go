package main

import (
	"encoding/binary"
	"fmt"
	"math"

	"github.com/h2so5/half"
)

func Float64frombytes(bytes []byte) float64 {
	bits := binary.LittleEndian.Uint64(bytes)
	float := math.Float64frombits(bits)
	return float
}

func Float64bytes(float float64) []byte {
	bits := math.Float64bits(float)
	bytes := make([]byte, 8)
	binary.LittleEndian.PutUint64(bytes, bits)
	return bytes
}

//ngetes ma
func Float16frombytes(bytes []byte) float32 {
	bits := binary.BigEndian.Uint32(bytes)

	floatx := math.Float32frombits(bits)

	float := half.NewFloat16(floatx)

	return float32(float)
}

func main() {
	//bytes := Float64bytes(math.Pi)
	//fmt.Println(bytes)
	//float := Float64frombytes(bytes)
	//fmt.Println(float)

	x := []byte{0x82, 0x71}

	f := binary.BigEndian.Uint16(x)
	sign := uint32((f >> 15) & 0x1)
	fmt.Println(sign)
	exp := (f >> 10) & 0x1f
	fmt.Println(exp)
	exp32 := uint32(exp) + 127 - 15
	if exp == 0 {
		exp32 = 0
	} else if exp == 0x1f {
		exp32 = 0xff
	}
	frac := uint32(f & 0x3ff)
	i := (sign << 31) | (exp32 << 23) | (frac << 13)
	b := math.Float32frombits(i)

	xz := []byte{0x82, 0x71, 0x00, 0x00}

	bz := Float16frombytes(xz)

	fmt.Println(b)
	fmt.Println(bz)
	fmt.Println(f)

}
