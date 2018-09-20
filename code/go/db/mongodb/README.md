# mongo1.go

參考： https://github.com/astaxie/build-web-application-with-golang/blob/master/zh/05.6.md

啟動 mongod

```
$ mongod --dbpath db

```

安裝 mgo

```
$go get gopkg.in/mgo.v2
```

執行範例 mongo1.db

```
$ go run mongo1.go
Phone: +55 53 8116 9639
```

結果執行成功 ....

## 帳號鎖定後

* https://blog.csdn.net/qq_38363459/article/details/80159387

```
func main() { 
dialInfo := &mgo.DialInfo{ 
Addrs: []string{"120.77.*.*:27017"}, //远程(或本地)服务器地址及端口号 
Direct: false, 
Timeout: time.Second * 1, 
Database: "admin", //数据库 
Source: "admin", 
Username: "root", 
Password: "root", 
PoolLimit: 4096, // Session.SetPoolLimit 
} 
session, err := mgo.DialWithInfo(dialInfo) 
if err != nil { 
panic(err) 
} 
defer session.Close() 

```