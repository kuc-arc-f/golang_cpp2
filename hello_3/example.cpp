#include <iostream>
#include <cstring>

extern "C" {
    // 1. GoLang から文字を 【送信】 される関数
    __declspec(dllexport) void SendStringToCpp(const char* str) {
        if (str != nullptr) {
            //std::cout << "[C++] Goから受信した文字列: " << str << std::endl;
            std::cout << "[C++] Go from Receive String: " << str << std::endl;
        }
    }

    // 2. GoLang へ文字を 【受信（返却）】 させる関数
    // Go側が用意したバッファ(buffer)に、C++の文字列を書き込みます
    __declspec(dllexport) int ReceiveStringFromCpp(char* buffer, int bufferSize) {
        const char* message = "Hello from C++ String!";
        int messageLen = std::strlen(message);

        // Go側のバッファサイズが足りているかチェック
        if (bufferSize > messageLen) {
            std::strcpy(buffer, message); // バッファにコピー
            return messageLen;            // コピーした文字数を返す
        }
        
        return -1; // バッファ不足エラー
    }
}