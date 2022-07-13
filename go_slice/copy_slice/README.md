# 指定具体函数,并指定cpu核数[因为是串行,提升核数可以看出对结果几乎无影响]
go test -bench='CopyArrNew' .
goos: linux
goarch: amd64
pkg: github.com/wsqyouth/coopers_go_code/go_slice/copy_slice
cpu: Intel(R) Xeon(R) Gold 6133 CPU @ 2.50GHz
BenchmarkCopyArrNew-2   	1000000000	         0.3696 ns/op

go test -bench='CopyArrNew' -cpu=1,2 .
goos: linux
goarch: amd64
pkg: github.com/wsqyouth/coopers_go_code/go_slice/copy_slice
cpu: Intel(R) Xeon(R) Gold 6133 CPU @ 2.50GHz
BenchmarkCopyArrNew     	1000000000	         0.3763 ns/op
BenchmarkCopyArrNew-2   	1000000000	         0.3695 ns/op

# 正则指定测试函数,默认执行1s, 表示用例执行1000000000次，每次操作约0.3693ns
go test -bench='CopyArr*' .
goos: linux
goarch: amd64
pkg: github.com/wsqyouth/coopers_go_code/go_slice/copy_slice
cpu: Intel(R) Xeon(R) Gold 6133 CPU @ 2.50GHz
BenchmarkCopyArr-2      	 5220757	       228.6 ns/op
BenchmarkCopyArrNew-2   	1000000000	         0.3693 ns/op
# 设置执行5s,进行3轮benchmark
go test -bench='CopyArr*' -benchtime=5s -count=3 .
goos: linux
goarch: amd64
pkg: github.com/wsqyouth/coopers_go_code/go_slice/copy_slice
cpu: Intel(R) Xeon(R) Gold 6133 CPU @ 2.50GHz
BenchmarkCopyArr-2      	26259784	       229.2 ns/op
BenchmarkCopyArr-2      	26426218	       227.8 ns/op
BenchmarkCopyArr-2      	26722513	       228.0 ns/op
BenchmarkCopyArrNew-2   	1000000000	         0.3662 ns/op
BenchmarkCopyArrNew-2   	1000000000	         0.3651 ns/op
BenchmarkCopyArrNew-2   	1000000000	         0.3672 ns/op

---
参考文档: https://geektutu.com/post/hpg-benchmark.html
