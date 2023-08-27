# try-generics-slices
Benchmarking the new generics package slices

```go
go test ./... -v -bench=. -benchmem -count 5                                                                   ─╯
goos: darwin
goarch: arm64
pkg: try-generics-slices
BenchmarkOneSort
BenchmarkOneSort-8     	 1000000	      1055 ns/op	     488 B/op	       3 allocs/op
BenchmarkOneSort-8     	 1000000	      1054 ns/op	     488 B/op	       3 allocs/op
BenchmarkOneSort-8     	 1000000	      1054 ns/op	     488 B/op	       3 allocs/op
BenchmarkOneSort-8     	 1000000	      1053 ns/op	     488 B/op	       3 allocs/op
BenchmarkOneSort-8     	 1000000	      1053 ns/op	     488 B/op	       3 allocs/op
BenchmarkTwoSlices
BenchmarkTwoSlices-8   	  186759	      6412 ns/op	       0 B/op	       0 allocs/op
BenchmarkTwoSlices-8   	  185872	      6418 ns/op	       0 B/op	       0 allocs/op
BenchmarkTwoSlices-8   	  186904	      6461 ns/op	       0 B/op	       0 allocs/op
BenchmarkTwoSlices-8   	  186676	      6508 ns/op	       0 B/op	       0 allocs/op
BenchmarkTwoSlices-8   	  181999	      6770 ns/op	       0 B/op	       0 allocs/op
PASS
ok  	try-generics-slices	13.974s
```

