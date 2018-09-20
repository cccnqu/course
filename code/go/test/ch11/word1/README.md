
```
PS D:\Dropbox\course\gopl.io\ch11\word1> go test
--- FAIL: TestFrenchPalindrome (0.00s)
        word_test.go:32: IsPalindrome("été") = false
--- FAIL: TestCanalPalindrome (0.00s)
        word_test.go:39: IsPalindrome("A man, a plan, a canal: Panama") = false
FAIL
exit status 1
FAIL    _/D_/Dropbox/course/gopl.io/ch11/word1  0.194s
PS D:\Dropbox\course\gopl.io\ch11\word1> go test -v
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
FAIL    _/D_/Dropbox/course/gopl.io/ch11/word1  0.162s

```