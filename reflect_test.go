package number

import (
	"reflect"
	"testing"
)

type myInt int

func TestPacked(t *testing.T) {
	cases := []any{
		int(0),
		int8(0),
		int16(0),
		int32(0),
		int64(0),
		uint(0),
		uint8(0),
		uint16(0),
		uint32(0),
		uint64(0),
		uintptr(0),
		float32(0),
		float64(0),
		complex64(0),
		complex128(0),
		myInt(0),
	}
	for _, c := range cases {
		c := c
		typ := reflect.TypeOf(c)
		t.Run(typ.Name(), func(t *testing.T) {
			v := types[typ.Kind()]
			if v == 0 {
				t.Fatalf("no Type")
			}
			if v.Bits() != typ.Bits() {
				t.Errorf("wrong bit size: want %d, got %d", typ.Bits(), v.Bits())
			}
			unsigned := reflect.Uint <= typ.Kind() && typ.Kind() <= reflect.Uintptr
			if v.Unsigned() != unsigned {
				t.Errorf("wrong unsignedness: want %t, got %t", unsigned, v.Unsigned())
			}
			fp := reflect.Float32 <= typ.Kind() && typ.Kind() <= reflect.Complex128
			if v.Float() != fp {
				t.Errorf("wrong floatingness: want %t, got %t", fp, v.Float())
			}
			ord := typ.Kind() != reflect.Complex128 && typ.Kind() != reflect.Complex64
			if v.Ordered() != ord {
				t.Errorf("wrong orderedness: want %t, got %t", ord, v.Ordered())
			}
		})
	}
}

func TestReflect(t *testing.T) {
	cases := []struct {
		t Type
		v any
	}{
		{Reflect[int](), int(0)},
		{Reflect[int8](), int8(0)},
		{Reflect[int16](), int16(0)},
		{Reflect[int32](), int32(0)},
		{Reflect[int64](), int64(0)},
		{Reflect[uint](), uint(0)},
		{Reflect[uint8](), uint8(0)},
		{Reflect[uint16](), uint16(0)},
		{Reflect[uint32](), uint32(0)},
		{Reflect[uint64](), uint64(0)},
		{Reflect[uintptr](), uintptr(0)},
		{Reflect[float32](), float32(0)},
		{Reflect[float64](), float64(0)},
		{Reflect[complex64](), complex64(0)},
		{Reflect[complex128](), complex128(0)},
		{Reflect[myInt](), myInt(0)},
	}
	for _, c := range cases {
		c := c
		typ := reflect.TypeOf(c.v)
		t.Run(typ.Name(), func(t *testing.T) {
			if c.t.packed != types[typ.Kind()] {
				t.Errorf("wrong type: want %b, got %b", c.t.packed, types[typ.Kind()])
			}
		})
	}
}

func TestReflectAlloc(t *testing.T) {
	var sink Type
	n := testing.AllocsPerRun(1, func() { sink = Reflect[complex128]() })
	if sink.packed != types[reflect.TypeOf(complex128(0)).Kind()] {
		t.Errorf("reflected wrong type! %#v", sink)
	}
	if n != 0 {
		t.Errorf("too many allocs: %v", n)
	}
}

func TestTypeOfAlloc(t *testing.T) {
	var sink Type
	n := testing.AllocsPerRun(1, func() { sink = TypeOf(complex128(0)) })
	if sink.packed != types[reflect.TypeOf(complex128(0)).Kind()] {
		t.Errorf("reflected wrong type! %#v", sink)
	}
	if n != 0 {
		t.Errorf("too many allocs: %v", n)
	}
}
