# Go Unit Test


```
PS D:\Dropbox\course\ws107a\go\test\01-unitTest> go test
--- FAIL: Test_Division_2 (0.00s)
        div_test.go:16: 就是不通过
FAIL
exit status 1
FAIL    _/D_/Dropbox/course/ws107a/go/test/01-unitTest  0.178s
PS D:\Dropbox\course\ws107a\go\test\01-unitTest> go test -v
=== RUN   Test_Division_1
--- PASS: Test_Division_1 (0.00s)
        div_test.go:11: 第一个测试通过了
=== RUN   Test_Division_2
--- FAIL: Test_Division_2 (0.00s)
        div_test.go:16: 就是不通过
FAIL
exit status 1
FAIL    _/D_/Dropbox/course/ws107a/go/test/01-unitTest  0.189s
```