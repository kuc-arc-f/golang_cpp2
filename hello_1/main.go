package main

import (
	"fmt"
	"syscall"
)

func main() {
	// 1. DLLをロード (遅延ロードが楽でおすすめ)
	dll := syscall.NewLazyDLL("example.dll")

	// 2. DLLから関数（プロシージャ）を取得
	addFunc := dll.NewProc("Add")
	helloFunc := dll.NewProc("HelloWorld")

	// 3. 引数がある関数を呼び出す
	// 引数はすべて uintptr にキャストして渡す必要があります
	a := 10
	b := 20
	r1, _, _ := addFunc.Call(uintptr(a), uintptr(b))
	
	fmt.Printf("C++からの計算結果: %d\n", int(r1))

	// 4. 引数がない関数を呼び出す
	helloFunc.Call()
}