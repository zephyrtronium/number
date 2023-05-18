# number [![Go Reference](https://pkg.go.dev/badge/github.com/zephyrtronium/number.svg)](https://pkg.go.dev/github.com/zephyrtronium/number)

Package number provides ultra-lightweight reflection for numeric types.

`number.Reflect[T]()` obtains a type descriptor for any numeric type `T` without reflection or allocation. It is about three times faster than package reflect to find the size of a type in bits:

```
$ go test -bench=.
goos: linux
goarch: amd64
pkg: github.com/zephyrtronium/number
cpu: Intel(R) Core(TM) i9-10885H CPU @ 2.40GHz
BenchmarkNumber-16              1000000000               0.9295 ns/op
BenchmarkNumberAny-16           1000000000               0.9223 ns/op
BenchmarkReflect-16             383663199                3.117 ns/op
PASS
ok      github.com/zephyrtronium/number 3.572s
```
