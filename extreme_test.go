package number_test

import (
	"fmt"
	"math"
	"testing"

	"github.com/zephyrtronium/number"
)

type myInt int

func TestGreatest(t *testing.T) {
	cases := []struct {
		got, want any
	}{
		{number.Greatest[int](), int(math.MaxInt)},
		{number.Greatest[int8](), int8(math.MaxInt8)},
		{number.Greatest[int16](), int16(math.MaxInt16)},
		{number.Greatest[int32](), int32(math.MaxInt32)},
		{number.Greatest[int64](), int64(math.MaxInt64)},
		{number.Greatest[uint](), uint(math.MaxUint)},
		{number.Greatest[uint8](), uint8(math.MaxUint8)},
		{number.Greatest[uint16](), uint16(math.MaxUint16)},
		{number.Greatest[uint32](), uint32(math.MaxUint32)},
		{number.Greatest[uint64](), uint64(math.MaxUint64)},
		{number.Greatest[uintptr](), ^uintptr(0)},
		{number.Greatest[float32](), float32(math.Inf(1))},
		{number.Greatest[float64](), math.Inf(1)},
		{number.Greatest[myInt](), myInt(math.MaxInt)},
	}
	for _, c := range cases {
		t.Run(fmt.Sprintf("%T", c.got), func(t *testing.T) {
			if c.want != c.got {
				t.Errorf("want != got: %#v != %#v", c.want, c.got)
			}
		})
	}
}

func TestGreatestFinite(t *testing.T) {
	cases := []struct {
		got, want any
	}{
		{number.GreatestFinite[int](), int(math.MaxInt)},
		{number.GreatestFinite[int8](), int8(math.MaxInt8)},
		{number.GreatestFinite[int16](), int16(math.MaxInt16)},
		{number.GreatestFinite[int32](), int32(math.MaxInt32)},
		{number.GreatestFinite[int64](), int64(math.MaxInt64)},
		{number.GreatestFinite[uint](), uint(math.MaxUint)},
		{number.GreatestFinite[uint8](), uint8(math.MaxUint8)},
		{number.GreatestFinite[uint16](), uint16(math.MaxUint16)},
		{number.GreatestFinite[uint32](), uint32(math.MaxUint32)},
		{number.GreatestFinite[uint64](), uint64(math.MaxUint64)},
		{number.GreatestFinite[uintptr](), ^uintptr(0)},
		{number.GreatestFinite[float32](), float32(math.MaxFloat32)},
		{number.GreatestFinite[float64](), float64(math.MaxFloat64)},
		{number.GreatestFinite[myInt](), myInt(math.MaxInt)},
	}
	for _, c := range cases {
		t.Run(fmt.Sprintf("%T", c.got), func(t *testing.T) {
			if c.want != c.got {
				t.Errorf("want != got: %#v != %#v", c.want, c.got)
			}
		})
	}
}

func TestLeast(t *testing.T) {
	cases := []struct {
		got, want any
	}{
		{number.Least[int](), int(math.MinInt)},
		{number.Least[int8](), int8(math.MinInt8)},
		{number.Least[int16](), int16(math.MinInt16)},
		{number.Least[int32](), int32(math.MinInt32)},
		{number.Least[int64](), int64(math.MinInt64)},
		{number.Least[uint](), uint(0)},
		{number.Least[uint8](), uint8(0)},
		{number.Least[uint16](), uint16(0)},
		{number.Least[uint32](), uint32(0)},
		{number.Least[uint64](), uint64(0)},
		{number.Least[uintptr](), uintptr(0)},
		{number.Least[float32](), float32(math.Inf(-1))},
		{number.Least[float64](), math.Inf(-1)},
		{number.Least[myInt](), myInt(math.MinInt)},
	}
	for _, c := range cases {
		t.Run(fmt.Sprintf("%T", c.got), func(t *testing.T) {
			if c.want != c.got {
				t.Errorf("want != got: %#v != %#v", c.want, c.got)
			}
		})
	}
}

func TestLeastFinite(t *testing.T) {
	cases := []struct {
		got, want any
	}{
		{number.LeastFinite[int](), int(math.MinInt)},
		{number.LeastFinite[int8](), int8(math.MinInt8)},
		{number.LeastFinite[int16](), int16(math.MinInt16)},
		{number.LeastFinite[int32](), int32(math.MinInt32)},
		{number.LeastFinite[int64](), int64(math.MinInt64)},
		{number.LeastFinite[uint](), uint(0)},
		{number.LeastFinite[uint8](), uint8(0)},
		{number.LeastFinite[uint16](), uint16(0)},
		{number.LeastFinite[uint32](), uint32(0)},
		{number.LeastFinite[uint64](), uint64(0)},
		{number.LeastFinite[uintptr](), uintptr(0)},
		{number.LeastFinite[float32](), float32(-math.MaxFloat32)},
		{number.LeastFinite[float64](), float64(-math.MaxFloat64)},
		{number.LeastFinite[myInt](), myInt(math.MinInt)},
	}
	for _, c := range cases {
		t.Run(fmt.Sprintf("%T", c.got), func(t *testing.T) {
			if c.want != c.got {
				t.Errorf("want != got: %#v != %#v", c.want, c.got)
			}
		})
	}
}

var (
	IntSink   int64
	FloatSink float64
)

func BenchmarkGreatestInt(b *testing.B) {
	for i := 0; i < b.N; i++ {
		IntSink = number.Greatest[int64]()
	}
}

func BenchmarkGreatestFloat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		FloatSink = number.Greatest[float64]()
	}
}
