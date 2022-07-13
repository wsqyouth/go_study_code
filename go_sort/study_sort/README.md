
[\u@dev \W]\go test -bench=. -benchmem
goos: linux
goarch: amd64
pkg: github.com/wsqyouth/coopers_go_code/go_sort/study_sort
cpu: Intel(R) Xeon(R) Gold 6133 CPU @ 2.50GHz
BenchmarkBubbleSort-2   	      26	  42210978 ns/op	       0 B/op	       0 allocs/op
BenchmarkSelectSort-2   	      25	  52937681 ns/op	       0 B/op	       0 allocs/op
BenchmarkQuickSort-2    	     100	  27571288 ns/op	       0 B/op	       0 allocs/op
PASS
ok  	github.com/wsqyouth/coopers_go_code/go_sort/study_sort	6.599s

分析：可以看到快排相对于冒泡有着一倍的提升
// https://mojotv.cn/algorithm/golang-quick-sort
