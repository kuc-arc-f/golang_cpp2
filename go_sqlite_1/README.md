# go_sqlite_1

 Version: 0.9.1

 date    : 2026/06/07

 update :

***

GoLang Window , C++ call TODO SQLite

* go version go1.25.3 windows/amd64
* LLMV CLang

***
### related

https://www.sqlite.org/download.html

* sqlite-amalgamation-*.zip , download
* sqlite3.h , sqlite3.c

***
* vcpkg install
```
.\vcpkg install nlohmann-json:x64-windows
```

***
* cpp build
```
.\build.bat
```

* build
```
go mod init example.com/go-sqlite-1
go mod tidy

go build main.go

```

***
* add
```
.\main.exe add hello
```

* list
```
.\main.exe list
```

* del
```
.\main.exe del 1
```




***

