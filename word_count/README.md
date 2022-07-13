
[\u@dev \W]\go test -bench=. -benchmem
goos: linux
goarch: amd64
pkg: github.com/wsqyouth/coopers_go_code/word_count
cpu: Intel(R) Xeon(R) Gold 6133 CPU @ 2.50GHz
BenchmarkGetAlphanumericNumByASCII-2    	44501674	        26.38 ns/op	       0 B/op	       0 allocs/op
BenchmarkGetAlphanumericNumByRegExp-2   	  290022	      3935 ns/op	    1925 B/op	      27 allocs/op

分析：可以看到使用正则设计到内存分配，性能也没有ascii方法好
参考链接：https://blog.csdn.net/K346K346/article/details/124936878
