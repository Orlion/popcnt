# popcnt
Improve popcnt performance by using simd instructions.

使用SIMD指令提升popcnt性能。

# Desc

对于一个`int64`数字来说可以直接使用`bits.OnesCount64`或者`popcnt`指令来获取其bit位为1的数量，但如果对于一个`[...]int64`这样的数组(比如bitmap)来说只能遍历这个数组挨个计算最后取和，性能比较差。

我们可以利用`SIMD`指令并行同时计算多个`int64`数字的bit位为1的个数，理论上可以提高N倍性能。

# Benchmark
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

可以看到使用SIMD指令优化之后性能下降了约150倍，与预期严重不符。

![](https://qn.doutub.com/1638175924196.jpg)

# 进一步优化思路

1. 从`popcnt.s`中汇编代码中可以看到逻辑较为啰嗦，还有优化空间。
2. 对于一个int64切片仍然要循环调用`SimdPopcntQuad`函数来计算，`SimdPopcntQuad`函数中又有许多寄存器初始化的代码，这样大量时间浪费寄存器初始化的工作上，可以改造下`SimdPopcntQuad`函数，使其接收一个`int64`切片，在一个汇编函数中直接计算出结果。