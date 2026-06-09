package main

import (
    "fmt"
    "syscall"
    "unsafe"
    "os"
    "strconv"
)

func testString() {
    // DLLの読み込み
    dll := syscall.NewLazyDLL("example.dll")
    sendStrFunc := dll.NewProc("SendStringToCpp")
    recvStrFunc := dll.NewProc("ReceiveStringFromCpp")

    // ==================================================
    // パターンA: GoLang から C++ に文字送信
    // ==================================================
    goMessage := "Hello from Go String!"
    // Goの文字列を、C言語形式（NULL終端のバイト配列）のポインタに変換
    cStrPointer, err := syscall.BytePtrFromString(goMessage)
    if err != nil {
        panic(err)
    }
    // C++の関数を呼び出し（ポインタをuintptrにキャスト）
    sendStrFunc.Call(uintptr(unsafe.Pointer(cStrPointer)))

    // ==================================================
    // パターンB: C++ から GoLang に文字受信
    // ==================================================
    // 1. Go側でデータを受け取るためのバッファ（スライス）をあらかじめ用意
    buffer := make([]byte, 256)    
    // 2. バッファの先頭ポインタと、そのサイズを取得
    bufPtr := uintptr(unsafe.Pointer(&buffer[0]))
    bufSize := uintptr(len(buffer))

    // 3. C++の関数を呼び出し、バッファに書き込んでもらう
    // 戻り値(r1)には、C++側で書き込まれた文字数が返ってきます
    r1, _, _ := recvStrFunc.Call(bufPtr, bufSize)
    stringLen := int(r1)

    if stringLen > 0 {
        // 4. 書き込まれたバイト数分だけ切り出して、Goの文字列に変換
        cppMessage := string(buffer[:stringLen])
        fmt.Printf("[Go]  C++から受信した文字列: %s\n", cppMessage)
    } else {
        fmt.Println("[Go]  文字列の受信に失敗しました（バッファサイズ不足など）")
    }
}

func main() {
    fmt.Println("全引数:", os.Args)
    if len(os.Args) < 2 {
        fmt.Println("error , argment none")
        return
    }    

    dll := syscall.NewLazyDLL("example.dll")
    todoAddFunc := dll.NewProc("todo_add")
    todoListFunc := dll.NewProc("todo_list")
    todoDeleteFunc := dll.NewProc("todo_delete")

    var argment = os.Args[1]
    if argment == "add" {
        if len(os.Args) < 3 {
            fmt.Println("error , argment none")
            return
        }
        var input = os.Args[2]
        fmt.Println("title=", input)
        title := input
        // Goの文字列を、C言語形式（NULL終端のバイト配列）のポインタに変換
        cStrPointer, err := syscall.BytePtrFromString(title)
        if err != nil {
            panic(err)
        }
        todoAddFunc.Call(uintptr(unsafe.Pointer(cStrPointer)))
    }
    if argment == "list" {
        todoListFunc.Call()
    }
    if argment == "del" {
        if len(os.Args) < 3 {
            fmt.Println("error , argment none")
            return
        }
        var id_str = os.Args[2]

        num, err := strconv.Atoi(id_str)
        if err != nil {
            fmt.Println("変換エラー:", err)
            return
        }
        fmt.Printf("型: %T, 値: %d\n", num, num) // 型: int, 値: 123        
        var id = num
        todoDeleteFunc.Call(uintptr(id))
    }
}
