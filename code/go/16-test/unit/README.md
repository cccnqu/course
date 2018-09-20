# 測試

參考: https://github.com/astaxie/build-web-application-with-golang/blob/master/zh/11.3.md


只報錯誤

```
PS D:\Dropbox\course\golang\example\16-test> go test
--- FAIL: Test_Division_2 (0.00s)
        unit_test.go:16: 就是不通过
FAIL
exit status 1
FAIL    _/D_/Dropbox/course/golang/example/16-test      0.214s
```

全部都報

```
PS D:\Dropbox\course\golang\example\16-test> go test
--- FAIL: Test_Division_2 (0.00s)
        unit_test.go:16: 就是不通过
FAIL
exit status 1
FAIL    _/D_/Dropbox/course/golang/example/16-test      0.214s
```

## cover 涵蓋度測試

```
PS D:\Dropbox\course\golang\example\16-test\gopl\coverage> go test -coverprofile=cover

sqrt(A / pi)
        map[A:87616 pi:3.141592653589793] => 167

pow(x, 3) + pow(y, 3)
        map[x:12 y:1] => 1729
        map[y:10 x:9] => 1729

5 / 9 * (F - 32)
        map[F:-40] => -40
        map[F:32] => 0
        map[F:212] => 100

-1 + -x
        map[x:1] => -2

-1 - x
        map[x:1] => -2
x % 2               unexpected '%'
math.Pi             unexpected '.'
!true               unexpected '!'
"hello"             unexpected '"'
log(10)             unknown function "log"
sqrt(1, 2)          call to sqrt has 2 args, want 1
PASS
coverage: 67.7% of statements
ok      _/D_/Dropbox/course/golang/example/16-test/gopl/coverage        0.264s
```
