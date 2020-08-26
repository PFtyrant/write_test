# write_test

```info
go test -v -run=BENCH -bench=. -benchmem
goos: linux
goarch: amd64
pkg: awesomeProject
BenchmarkPrintln
BenchmarkPrintln-4                 87802             13421 ns/op              80 B/op          5 allocs/op
BenchmarkConcat
BenchmarkConcat-4                 377556              3178 ns/op              48 B/op          6 allocs/op
BenchmarkConcatOneLine
BenchmarkConcatOneLine-4          364741              2960 ns/op              32 B/op          2 allocs/op
PASS
ok      awesomeProject  4.678s
```
