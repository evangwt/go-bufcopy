# go-bufcopy
golang copy io stream optimised by using sync.Pool

# Benchmark
```
go test -bench . -benchtime=50000x
goos: darwin
goarch: amd64
pkg: github.com/evangwt/go-bufcopy
BenchmarkIoCopySmall-12     	   50000	       738 ns/op	    1342 B/op	       0 allocs/op
BenchmarkBufCopySmall-12    	   50000	       393 ns/op	    1342 B/op	       0 allocs/op
BenchmarkIoCopyLarge-12     	   50000	     33851 ns/op	   85887 B/op	       0 allocs/op
BenchmarkBufCopyLarge-12    	   50000	     13440 ns/op	   85887 B/op	       0 allocs/op
PASS
ok  	github.com/evangwt/go-bufcopy	12.514s
```
