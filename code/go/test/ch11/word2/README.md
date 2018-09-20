
```
PS D:\Dropbox\course\gopl.io\ch11\word2> go test -bench=BenchmarkIsPalindrome
goos: windows
goarch: amd64
BenchmarkIsPalindrome-4          1000000              1665 ns/op
PASS
ok      _/D_/Dropbox/course/gopl.io/ch11/word2  1.887s
```


```
$ go test -cpuprofile=cpu.out 
$ go test -blockprofile=block.out 
$ go test -memprofile=mem.out
```