# countServer.go

啟動 server

```
PS D:\Dropbox\course\ws107a\go\server> go run countServer.go

```

執行 client

```
PS D:\Dropbox\course\ws107a\go\client\fetch> ./fetch http://localhost:8000/help
URL.path = "/help"
PS D:\Dropbox\course\ws107a\go\client\fetch> ./fetch http://localhost:8000/help
URL.path = "/help"
PS D:\Dropbox\course\ws107a\go\client\fetch> ./fetch http://localhost:8000
URL.path = "/"
PS D:\Dropbox\course\ws107a\go\client\fetch> ./fetch http://localhost:8000/count
Count 3
PS D:\Dropbox\course\ws107a\go\client\fetch> ./fetch http://localhost:8000/count
Count 3
PS D:\Dropbox\course\ws107a\go\client\fetch> ./fetch http://localhost:8000/count
Count 3
PS D:\Dropbox\course\ws107a\go\client\fetch> ./fetch http://localhost:8000
URL.path = "/"
PS D:\Dropbox\course\ws107a\go\client\fetch> ./fetch http://localhost:8000/count
Count 4
```

