# Go 測試

## word1

```
PS D:\Dropbox\course\golang\example\16-test\gopl> cd word1
PS D:\Dropbox\course\golang\example\16-test\gopl\word1> go test
--- FAIL: TestFrenchPalindrome (0.00s)
        word_test.go:32: IsPalindrome("été") = false
--- FAIL: TestCanalPalindrome (0.00s)
        word_test.go:39: IsPalindrome("A man, a plan, a canal: Panama") = false
FAIL
exit status 1
FAIL    _/D_/Dropbox/course/golang/example/16-test/gopl/word1   0.166s
PS D:\Dropbox\course\golang\example\16-test\gopl\word1> go test -v
=== RUN   TestPalindrome
--- PASS: TestPalindrome (0.00s)
=== RUN   TestNonPalindrome
--- PASS: TestNonPalindrome (0.00s)
=== RUN   TestFrenchPalindrome
--- FAIL: TestFrenchPalindrome (0.00s)
        word_test.go:32: IsPalindrome("été") = false
=== RUN   TestCanalPalindrome
--- FAIL: TestCanalPalindrome (0.00s)
        word_test.go:39: IsPalindrome("A man, a plan, a canal: Panama") = false
FAIL
exit status 1
FAIL    _/D_/Dropbox/course/golang/example/16-test/gopl/word1   0.225s
```

## word2

```
PS D:\Dropbox\course\golang\example\16-test\gopl\word2> go test
PASS
ok      _/D_/Dropbox/course/golang/example/16-test/gopl/word2   0.191s
PS D:\Dropbox\course\golang\example\16-test\gopl\word2> go test -v
=== RUN   TestIsPalindrome
--- PASS: TestIsPalindrome (0.00s)
=== RUN   TestRandomPalindromes
--- PASS: TestRandomPalindromes (0.00s)
        word_test.go:92: Random seed: 1529385576805096900
=== RUN   ExampleIsPalindrome
--- PASS: ExampleIsPalindrome (0.00s)
PASS
ok      _/D_/Dropbox/course/golang/example/16-test/gopl/word2   0.195s
```

BenchMark

```
PS D:\Dropbox\course\golang\example\16-test\gopl\word2> go test -bench=.
PASS
ok      _/D_/Dropbox/course/golang/example/16-test/gopl/word2   0.199s
```

## echo


```


