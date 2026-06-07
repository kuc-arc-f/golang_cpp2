# hello_1

 Version: 0.9.1

 date    : 2026/06/07

 update :

***

GoLang Window , C++ call example

***
* go version go1.25.3 windows/amd64
* gcc

***
* cpp build
```
g++ -shared -o example.dll example.cpp 
```

* build
```
go mod init example.com/cgo-hello-1
go mod tidy

go build main.go
```


***

