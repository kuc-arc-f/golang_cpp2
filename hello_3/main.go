package main

import (
	"fmt"
	"syscall"
	"unsafe"
)

func main() {
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