# go-bufcopy
golang copy io stream optimised by using sync.Pool

# Benchmark
```
go test -bench .
goos: darwin
goarch: amd64
pkg: bufcopy
BenchmarkIoCopySmall-12     	 2748542	       476 ns/op	    1562 B/op	       0 allocs/op
BenchmarkBufCopySmall-12    	 8255025	       345 ns/op	    1040 B/op	       0 allocs/op
BenchmarkIoCopyLarge-12     	   80940	     13775 ns/op	  106119 B/op	       0 allocs/op
BenchmarkBufCopyLarge-12    	  126852	      8050 ns/op	   67711 B/op	       0 allocs/op
PASS
ok  	bufcopy	8.500s
```
