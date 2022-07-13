
go test -bench .

[\u@dev \W]\go test -bench='RemoveDuplicate' -benchtime=5s -count=3 .
goos: linux
goarch: amd64
pkg: github.com/wsqyouth/coopers_go_code/go_slice/remove_duplicate
cpu: Intel(R) Xeon(R) Gold 6133 CPU @ 2.50GHz
BenchmarkRemoveDuplicateElement-2   	19853739	       296.9 ns/op
BenchmarkRemoveDuplicateElement-2   	19740693	       296.6 ns/op
BenchmarkRemoveDuplicateElement-2   	20175840	       297.3 ns/op
BenchmarkRemoveDuplicateString-2    	10431752	       572.9 ns/op
BenchmarkRemoveDuplicateString-2    	10527133	       581.2 ns/op
BenchmarkRemoveDuplicateString-2    	10439780	       574.7 ns/op
