# Go

* Go With MongoDB 2 -- https://www.jianshu.com/p/6e196b5f0113
* http://golang-zhtw.netdpi.net/
* 有效地去 - Effective go 正體中文翻譯 -- https://ronmi.github.io/post/go/effectivego/
* https://stackoverflow.com/questions/29898400/import-struct-from-another-package-and-file-golang

## 自製套件

目前的做法

把 shop.go 放在 $GOPATH/src/misavo 底下 (C:\mygo\src\misavo )，然後用切到該資料夾後，用 go install ，
此時 C:\mygo\pkg\windows_amd64\ 底下會產生一個 misavo.a 的套件，然後就可以在程式裡用 import "misavo" ，
然後用 misavo.函數 或 misavo.結構 來使用了。

## 指令

```
$ go env GOPATH
c:\mygo

$ go env
set GOARCH=amd64
set GOBIN=
set GOCACHE=C:\Users\user\AppData\Local\go-build
set GOEXE=.exe
set GOHOSTARCH=amd64
set GOHOSTOS=windows
set GOOS=windows
set GOPATH=c:\mygo
set GORACE=
set GOROOT=C:\Go
set GOTMPDIR=
set GOTOOLDIR=C:\Go\pkg\tool\windows_amd64
set GCCGO=gccgo
set CC=gcc
set CXX=g++
set CGO_ENABLED=1
set CGO_CFLAGS=-g -O2
set CGO_CPPFLAGS=
set CGO_CXXFLAGS=-g -O2
set CGO_FFLAGS=-g -O2
set CGO_LDFLAGS=-g -O2
set PKG_CONFIG=pkg-config
set GOGCCFLAGS=-m64 -mthreads -fmessage-length=0 -fdebug-prefix-map=C:\Users\user\AppData\Local\Temp\go-build941532770=/tmp/go-build -gno-record-gcc-switches
```