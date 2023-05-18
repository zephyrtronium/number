package number_test

import (
	"fmt"

	"github.com/zephyrtronium/number"
)

func Min[T ~int32 | ~float32](xs ...T) T {
	min := number.Greatest[T]()
	for _, x := range xs {
		if x < min {
			min = x
		}
	}
	return min
}

func ExampleGreatest_sliceMin() {
	fmt.Println(Min[int32](1, 3, 2))
	fmt.Println(Min[int32]())
	type myFloat float32
	fmt.Println(Min[myFloat](1.5))
	// Output:
	// 1
	// 2147483647
	// 1.5
}
