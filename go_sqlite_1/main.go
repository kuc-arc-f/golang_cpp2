package main

import (
	"fmt"
	"syscall"
	"unsafe"
    "os"
    "strconv"
)

func test(){
    // 1. DLLのロード
    // Windows環境で実行時にカレントディレクトリにある "mydll.dll" をロードします
    dll, err := syscall.LoadDLL("mydll.dll")
    if err != nil {
        fmt.Printf("DLLのロードに失敗しました: %v\n", err)
        fmt.Println("ビルドした mydll.dll が Go 実行バイナリと同じディレクトリにあることを確認してください。")
        return
    }
    defer dll.Release()

    fmt.Println("==================================================")
    fmt.Println(" 1. GoからC++ DLLへ文字列を送信する")
    fmt.Println("==================================================")
    sendToCPP, err := dll.FindProc("SendToCPP")
    if err != nil {
        fmt.Printf("SendToCPP 関数の検出に失敗しました: %v\n", err)
        return
    }
    messageToGo := "Hello from GoLang (Windows) -123"
    // Goの文字列をヌル終端のC文字列バイトポインタに変換
    cStr, err := syscall.BytePtrFromString(messageToGo)
    if err != nil {
        fmt.Printf("文字列の変換に失敗しました: %v\n", err)
        return
    }

    // DLL関数の呼び出し (uintptrにキャストしてポインタを渡す)
    _, _, _ = sendToCPP.Call(uintptr(unsafe.Pointer(cStr)))
    fmt.Printf("Go側で送信完了: \"%s\"\n\n", messageToGo)

    fmt.Println("==================================================")
    fmt.Println(" 3. C++ DLLから文字列を受信する (ポインタ返却方式)")
    fmt.Println("==================================================")
    getStringFromCPP, err := dll.FindProc("GetStringFromCPP")
    if err != nil {
        fmt.Printf("GetStringFromCPP 関数の検出に失敗しました: %v\n", err)
        return
    }
    freeCPPString, err := dll.FindProc("FreeCPPString")
    if err != nil {
        fmt.Printf("FreeCPPString 関数の検出に失敗しました: %v\n", err)
        return
    }

    // C++側でヒープメモリ上に確保した文字列のポインタを受け取る
    rptr, _, _ := getStringFromCPP.Call()
    if rptr != 0 {
        // ヌル終端までの長さを算出する
        length := 0
        for {
            val := *(*byte)(unsafe.Pointer(rptr + uintptr(length)))
            if val == 0 {
                break
            }
            length++
        }

        // ポインタからGoのバイトスライスを作成して文字列に変換
        // Go 1.20以降の推奨される書き方 `unsafe.Slice` を使用します
        ptr := (*byte)(unsafe.Pointer(rptr))
        slice := unsafe.Slice(ptr, length)
        receivedStr2 := string(slice)

        fmt.Printf("C++から受信した文字列: \"%s\"\n", receivedStr2)

        // C++側で確保したメモリを解放するためにFree関数を呼び出します
        _, _, _ = freeCPPString.Call(rptr)
        fmt.Println("C++側のメモリを正常に解放しました。")
    } else {
        fmt.Println("ポインタの取得に失敗しました。")
    }
    fmt.Println("==================================================")
}

/**
*
* @param
*
* @return
*/
func main() {
    fmt.Println("全引数:", os.Args)
    if len(os.Args) < 2 {
        fmt.Println("error , argment none")
        return
    } 
    dll, err := syscall.LoadDLL("mydll.dll")
    if err != nil {
        fmt.Printf("DLLのロードに失敗しました: %v\n", err)
        fmt.Println("ビルドした mydll.dll が Go 実行バイナリと同じディレクトリにあることを確認してください。")
        return
    }
    defer dll.Release()

    var argment = os.Args[1]
    if argment == "add" {
        if len(os.Args) < 3 {
            fmt.Println("error , argment none")
            return
        }
        var input = os.Args[2]
        fmt.Println("title=", input)
        title := input

        todoAdd, err := dll.FindProc("TodoAdd")
        if err != nil {
            fmt.Printf("TodoAdd 関数の検出に失敗しました: %v\n", err)
            return
        }
        // Goの文字列をヌル終端のC文字列バイトポインタに変換
        cStr, err := syscall.BytePtrFromString(title)
        if err != nil {
            fmt.Printf("文字列の変換に失敗しました: %v\n", err)
            return
        }
        // DLL関数の呼び出し (uintptrにキャストしてポインタを渡す)
        _, _, _ = todoAdd.Call(uintptr(unsafe.Pointer(cStr)))
        if err != nil {
            fmt.Printf("TodoAdd 関数の検出に失敗しました: %v\n", err)
            return
        }
    }       
    if argment == "list" {
        todoList, err := dll.FindProc("TodoList")
        if err != nil {
            fmt.Printf("TodoList 関数の検出に失敗しました: %v\n", err)
            return
        }
        freeCPPString, err := dll.FindProc("FreeCPPString")
        if err != nil {
            fmt.Printf("FreeCPPString 関数の検出に失敗しました: %v\n", err)
            return
        }        
        rptr, _, _ := todoList.Call()
        if rptr != 0 {
            // ヌル終端までの長さを算出する
            length := 0
            for {
                val := *(*byte)(unsafe.Pointer(rptr + uintptr(length)))
                if val == 0 {
                    break
                }
                length++
            }
            ptr := (*byte)(unsafe.Pointer(rptr))
            slice := unsafe.Slice(ptr, length)
            receivedStr2 := string(slice)

            fmt.Printf("C++から受信した文字列: \"%s\"\n", receivedStr2)

            // C++側で確保したメモリを解放するためにFree関数を呼び出します
            _, _, _ = freeCPPString.Call(rptr)
            fmt.Println("C++側のメモリを正常に解放しました。")
        } else {
            fmt.Println("ポインタの取得に失敗しました。")
        }        
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
        //TodoDelete
        todoDelete, err := dll.FindProc("TodoDelete")
        if err != nil {
            fmt.Printf("TodoDelete 関数の検出に失敗しました: %v\n", err)
            return
        }
        // DLL関数の呼び出し (uintptrにキャストしてポインタを渡す)
        _, _, _ = todoDelete.Call(uintptr(id))
        if err != nil {
            fmt.Printf("TodoDelete 関数の検出に失敗しました: %v\n", err)
            return
        }        

    }

}
