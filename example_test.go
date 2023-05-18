package number_test

import (
	"fmt"

	"github.com/zephyrtronium/number"
)

func print16bitInfo[T ~int16 | ~uint16](name string) {
	t := number.Reflect[T]()
	fmt.Printf("%s is %d bits, unsigned: %t, float: %t, ordered: %t\n", name, t.Bits(), t.Unsigned(), t.Float(), t.Ordered())
}

func Example() {
	type myInt int16
	print16bitInfo[uint16]("uint16")
	print16bitInfo[myInt]("myInt")
	// Output:
	// uint16 is 16 bits, unsigned: true, float: false, ordered: true
	// myInt is 16 bits, unsigned: false, float: false, ordered: true
}
