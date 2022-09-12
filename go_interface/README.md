
[\u@dev \W]\go test -bench=. -benchmem
goos: linux
goarch: amd64
pkg: github.com/wsqyouth/coopers_go_code/go_interface
cpu: Intel(R) Xeon(R) Gold 6133 CPU @ 2.50GHz
BenchmarkCheckParamSample-2   	  712669	      1482 ns/op	     360 B/op	       5 allocs/op
BenchmarkCheckParam-2         	  596887	      1720 ns/op	     352 B/op	       5 allocs/op
PASS
思考：比较了使用反射和简单值的校验，可以发现二者的性能没差太多

