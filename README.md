# popcnt
Improve popcnt performance by using simd instructions

使用SIMD指令优化popcnt性能

# benchmark
```
# go test -bench=. -v
=== RUN   TestSimdPopcntQuad
--- PASS: TestSimdPopcntQuad (0.00s)
goos: linux
goarch: amd64
pkg: github.com/Orlion/popcnt
cpu: Intel Core Processor (Broadwell, no TSX)
BenchmarkSimdPopcntQuad
BenchmarkSimdPopcntQuad-8        3693530               330.8 ns/op
BenchmarkSerial
BenchmarkSerial-8               539924296                2.232 ns/op
PASS
ok      github.com/Orlion/popcnt        2.993s
```

可以看到使用SIMD指令优化之后性能下降了约150倍

![](https://qn.doutub.com/1638175924196.jpg)

# 放弃了