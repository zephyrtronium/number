// Package number provides lightweight reflection on numeric types.
package number

import "unsafe"

// Type is metadata about a numeric type.
type Type struct {
	packed
}

// Reflect gets a Type for a numeric type.
func Reflect[T ~int | ~int8 | ~int16 | ~int32 | ~int64 | ~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr | ~float32 | ~float64 | ~complex64 | ~complex128]() Type {
	return Type{types[ekind(T(0))]}
}

// TypeOf gets a Type for a numeric value.
func TypeOf[T ~int | ~int8 | ~int16 | ~int32 | ~int64 | ~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr | ~float32 | ~float64 | ~complex64 | ~complex128](x T) Type {
	return Type{types[ekind(x)]}
}

// TypeOfAny gets a Type for an arbitrary value. Panics if the dynamic type of
// x is not numeric.
func TypeOfAny(x any) Type {
	r := types[ekind(x)]
	if r == 0 {
		panic("number: non-numeric type")
	}
	return Type{r}
}

// ekind pulls the reflect.Kind from an empty interface directly.
func ekind(x any) uint8 {
	// Partial definition of internal/abi.Type, up to the reflect.Kind.
	type rtype struct {
		_    uintptr // size
		_    uintptr // ptrbytes
		_    uint32  // hash
		_    uint8   // tflag
		_    uint8   // align
		_    uint8   // fieldalign
		kind uint8
		// equal func(pointer, pointer) bool
		// gcdata *byte
		// str int32
		// ptr int32
	}
	type eface struct {
		rtype *rtype
		// val unsafe.Pointer
	}
	e := (*eface)(unsafe.Pointer(&x))
	return e.rtype.kind
}

// packed is numeric type metadata as packed bitfields.
// Bits 0..3 map the size of the type: 001 => 8 bits, 010 => 16, 011 => 32,
// 100 => 64, 101 => 128; other values invalid.
// Bit 3 is set if the type is unsigned.
// Bit 4 is set if the type is floating point.
// Bit 5 is set if the type is ordered.
type packed uint8

var types = [256]packed{
	2:  0b1_0_0_000 | intsize, // int
	3:  0b1_0_0_001,           // int8
	4:  0b1_0_0_010,           // int16
	5:  0b1_0_0_011,           // int32
	6:  0b1_0_0_100,           // int64
	7:  0b1_0_1_000 | intsize, // uint
	8:  0b1_0_1_001,           // uint8
	9:  0b1_0_1_010,           // uint16
	10: 0b1_0_1_011,           // uint32
	11: 0b1_0_1_100,           // uint64
	12: 0b1_0_1_000 | ptrsize, // uintptr
	13: 0b1_1_0_011,           // float32
	14: 0b1_1_0_100,           // float64
	15: 0b0_1_0_100,           // complex64
	16: 0b0_1_0_101,           // complex128
}

const (
	intsize = 0b011 + packed(^uint(0)>>32&1)
	ptrsize = 0b011 + packed(^uintptr(0)>>32&1)

	sizeflags  = 0b0_0_0_111
	unsignflag = 0b0_0_1_000
	fpflag     = 0b0_1_0_000
	ordflag    = 0b1_0_0_000
)

var bitmap = [8]uint8{
	0b001: 8,
	0b010: 16,
	0b011: 32,
	0b100: 64,
	0b101: 128,
}

// Bits returns the size of the type in bits.
func (p packed) Bits() int {
	return int(bitmap[p&sizeflags])
}

// Unsigned returns whether the type is unsigned.
func (p packed) Unsigned() bool {
	return p&unsignflag != 0
}

// Float returns whether the type is floating-point.
func (p packed) Float() bool {
	return p&fpflag != 0
}

// Ordered returns whether the type is ordered.
func (p packed) Ordered() bool {
	return p&ordflag != 0
}
