#ifndef DLL_H
#define DLL_H

// Windows DLLのエクスポートマクロ定義
#ifdef BUILD_DLL
    #define DLL_EXPORT __declspec(dllexport)
#else
    #define DLL_EXPORT __declspec(dllimport)
#endif

extern "C" {
    // ポインタ返却方式で確保されたメモリを解放する関数
    DLL_EXPORT void FreeCPPString(const char* ptr);

    DLL_EXPORT void TodoAdd(const char* str);

    DLL_EXPORT void TodoDelete(int id);

    DLL_EXPORT const char* TodoList();

    // 1. GoからC++へ文字列を送信する関数
    DLL_EXPORT void SendToCPP(const char* str);

    // 2. C++からGoへ文字列を受信する関数 (Go側で用意したバッファに書き込む方式)
    DLL_EXPORT int ReceiveFromCPP(char* buffer, int maxLen);

    // 3. C++からGoへ文字列を受信する関数 (C++側で確保したメモリのポインタを返す方式)
    DLL_EXPORT const char* GetStringFromCPP();

}

#endif // DLL_H
