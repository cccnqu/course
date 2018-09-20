# Go 與 Linux

* [Google 為什麼要從頭開發一套新的操作系統 Fuchsia ？](https://www.inside.com.tw/2016/08/22/google-fuchsia)
* https://github.com/astaxie/go-system-programming (不太好)
  * 中文版 -- https://github.com/astaxie/go-system-programming/tree/master/zh
* https://golang.org/pkg/syscall/
* https://github.com/coreos/pkg/blob/master/dlopen/dlopen_example.go

## Linux

* [How to Install Go 1.6 on Ubuntu 16.04](https://www.digitalocean.com/community/tutorials/how-to-install-go-1-6-on-ubuntu-16-04)

```
$ sudo apt-get install golang-go // 在我電腦上是安裝 go-1.6
...

root@ubuntu-512mb-sgp1-01:~/ccc/go# ls /usr/lib/go-1.6
bin  pkg  src  test  VERSION
root@ubuntu-512mb-sgp1-01:~/ccc/go# ls /usr/lib/go-1.6/bin
go  gofmt
root@ubuntu-512mb-sgp1-01:~/ccc/go# ls /usr/lib/go-1.6/bin

```

要安裝更新版，請用下列方式：

* [How to: Install Go 1.8 on Ubuntu 16.04](https://medium.com/@patdhlk/how-to-install-go-1-8-on-ubuntu-16-04-710967aa53c9)

```
root@ubuntu-512mb-sgp1-01:~/ccc/go# go version
-bash: /usr/bin/go: No such file or directory
root@ubuntu-512mb-sgp1-01:~/ccc/go# sudo vim /root/.profile
root@ubuntu-512mb-sgp1-01:~/ccc/go# sudo vim /root/.bashrc
root@ubuntu-512mb-sgp1-01:~/ccc/go# go version
-bash: /usr/bin/go: No such file or directory
root@ubuntu-512mb-sgp1-01:~/ccc/go# sudo vim /root/.profile
root@ubuntu-512mb-sgp1-01:~/ccc/go# echo $PATH
/usr/local/sbin:/usr/local/bin:/usr/sbin:/usr/bin:/sbin:/bin:/usr/games:/usr/local/games:/snap/bin
root@ubuntu-512mb-sgp1-01:~/ccc/go# sudo vim /root/.profile
root@ubuntu-512mb-sgp1-01:~/ccc/go# source ~/.profile
root@ubuntu-512mb-sgp1-01:~/ccc/go# echo $PATH
/usr/local/sbin:/usr/local/bin:/usr/sbin:/usr/bin:/sbin:/bin:/usr/games:/usr/local/games:/snap/bin:/usr/local/go/bin
root@ubuntu-512mb-sgp1-01:~/ccc/go# go
Go is a tool for managing Go source code.

Usage:

        go command [arguments]
```

## 設定 GOPATH

```
root@ubuntu-512mb-sgp1-01:~/ccc/go# vim ~/.profile
root@ubuntu-512mb-sgp1-01:~/ccc/go# source ~/.profile
root@ubuntu-512mb-sgp1-01:~/ccc/go# echo $GOPATH
/root/ccc/go

```

## Go 和 WebAssembly

* meta: WebAssembly ("wasm") support  -- https://github.com/golang/go/issues/18892
  * 目前看來，短期內 go 沒有支持 webassembly 的計畫