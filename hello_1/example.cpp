#include <iostream>

// extern "C" でC言語互換にし、__declspec(dllexport) でDLL外から呼べるようにする
extern "C" {
    __declspec(dllexport) int Add(int a, int b) {
        return a + b;
    }

    __declspec(dllexport) void HelloWorld() {
        std::cout << "Hello from C++!" << std::endl;
    }
}