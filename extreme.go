package number

// Greatest returns the value such that T(x) <= Greatest[T]() is true for all
// non-NaN x.
// In particular, this returns +Inf for float32 and float64. Use
// [GreatestFinite] instead if you want finite values.
func Greatest[T ~int | ~int8 | ~int16 | ~int32 | ~int64 | ~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr | ~float32 | ~float64]() T {
	typ := Reflect[T]()
	switch typ.packed {
	case 0b1_0_0_001: // int8
		return T(127)
	case 0b1_0_0_010: // int16
		var x int16 = 0x7fff
		return T(x)
	case 0b1_0_0_011: // int32 or int
		var x int32 = 0x7fffffff
		return T(x)
	case 0b1_0_0_100: // int64 or int
		var x int64 = 0x7fffffffffffffff
		return T(x)
	case 0b1_0_1_001: // uint8
		var x uint8 = 0xff
		return T(x)
	case 0b1_0_1_010: // uint16
		var x uint16 = 0xffff
		return T(x)
	case 0b1_0_1_011: // uint32, uint, uintptr
		var x uint32 = 0xffffffff
		return T(x)
	case 0b1_0_1_100: // uint64, uint, uintptr
		var x uint64 = 0xffffffffffffffff
		return T(x)
	case 0b1_1_0_011: // float32
		var x float32 = 0x1p127
		return T(2 * x)
	case 0b1_1_0_100: // float64
		var x float64 = 0x1p1023
		return T(2 * x)
	default:
		panic("unreachable")
	}
}

// GreatestFinite returns the largest finite value of a numeric type.
func GreatestFinite[T ~int | ~int8 | ~int16 | ~int32 | ~int64 | ~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr | ~float32 | ~float64]() T {
	typ := Reflect[T]()
	switch typ.packed {
	case 0b1_0_0_001: // int8
		return T(127)
	case 0b1_0_0_010: // int16
		var x int16 = 0x7fff
		return T(x)
	case 0b1_0_0_011: // int32 or int
		var x int32 = 0x7fffffff
		return T(x)
	case 0b1_0_0_100: // int64 or int
		var x int64 = 0x7fffffffffffffff
		return T(x)
	case 0b1_0_1_001: // uint8
		var x uint8 = 0xff
		return T(x)
	case 0b1_0_1_010: // uint16
		var x uint16 = 0xffff
		return T(x)
	case 0b1_0_1_011: // uint32, uint, uintptr
		var x uint32 = 0xffffffff
		return T(x)
	case 0b1_0_1_100: // uint64, uint, uintptr
		var x uint64 = 0xffffffffffffffff
		return T(x)
	case 0b1_1_0_011: // float32
		var x float32 = 0x1p127 * (1 + (1 - 0x1p-23))
		return T(x)
	case 0b1_1_0_100: // float64
		var x float64 = 0x1p1023 * (1 + (1 - 0x1p-52))
		return T(x)
	default:
		panic("unreachable")
	}
}

// Least returns the value such that Least[T]() <= T(x) is true for all x.
// In particular, this returns -Inf for float32 and float64. Use [LeastFinite]
// instead if you want finite values.
func Least[T ~int | ~int8 | ~int16 | ~int32 | ~int64 | ~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr | ~float32 | ~float64]() T {
	typ := Reflect[T]()
	switch typ.packed {
	case 0b1_0_0_001: // int8
		var x int8 = -128
		return T(x)
	case 0b1_0_0_010: // int16
		var x int16 = -0x8000
		return T(x)
	case 0b1_0_0_011: // int32 or int
		var x int32 = -0x80000000
		return T(x)
	case 0b1_0_0_100: // int64 or int
		var x int64 = -0x8000000000000000
		return T(x)
	case 0b1_0_1_001, 0b1_0_1_010, 0b1_0_1_011, 0b1_0_1_100: // all uints
		return 0
	case 0b1_1_0_011: // float32
		var x float32 = -0x1p127
		return T(2 * x)
	case 0b1_1_0_100: // float64
		var x float64 = -0x1p1023
		return T(2 * x)
	default:
		panic("unreachable")
	}
}

// LeastFinite returns the least finite value of a numeric type.
func LeastFinite[T ~int | ~int8 | ~int16 | ~int32 | ~int64 | ~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr | ~float32 | ~float64]() T {
	typ := Reflect[T]()
	switch typ.packed {
	case 0b1_0_0_001: // int8
		var x int8 = -128
		return T(x)
	case 0b1_0_0_010: // int16
		var x int16 = -0x8000
		return T(x)
	case 0b1_0_0_011: // int32 or int
		var x int32 = -0x80000000
		return T(x)
	case 0b1_0_0_100: // int64 or int
		var x int64 = -0x8000000000000000
		return T(x)
	case 0b1_0_1_001, 0b1_0_1_010, 0b1_0_1_011, 0b1_0_1_100: // all uints
		return 0
	case 0b1_1_0_011: // float32
		var x float32 = -0x1p127 * (1 + (1 - 0x1p-23))
		return T(x)
	case 0b1_1_0_100: // float64
		var x float64 = -0x1p1023 * (1 + (1 - 0x1p-52))
		return T(x)
	default:
		panic("unreachable")
	}
}
