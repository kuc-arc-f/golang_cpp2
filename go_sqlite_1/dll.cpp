#define BUILD_DLL
#include "dll.h"
#include <iostream>
#include <string>
#include <cstring>
#include <algorithm>

#include "include/my_todo.hpp"

std::string STR_BUFF = "";

// ポインタ返却方式で確保されたメモリを解放する関数
void FreeCPPString(const char* ptr) {
    if (ptr != nullptr) {
        delete[] ptr;
        // std::cout << "[C++ DLL] Memory freed successfully." << std::endl;
    }
}

void TodoAdd(const char* str) {
    if (str != nullptr) {
        MyTodo todo_helper("");
        todo_helper.todo_add_handler(std::string(str));    
        std::cout << "[C++] Go from Receive String: " << str << std::endl;
    }
}

const char* TodoList() {
    MyTodo todo_helper("");
    std::string response  = todo_helper.todo_list_handler(); 
    STR_BUFF = response;   

    char* buffer = new char[response.length() + 1];
    std::memcpy(buffer, response.c_str(), response.length());
    buffer[response.length()] = '\0';
    
    return buffer;
}

void TodoDelete(int id) {
    MyTodo todo_helper("");
    todo_helper.todo_delete_handler(id);    
}

// 1. GoからC++へ文字列を送信する関数
void SendToCPP(const char* str) {
    if (str != nullptr) {
        std::cout << "[C++ DLL] Received string from Go: \"" << str << "\"" << std::endl;
    } else {
        std::cout << "[C++ DLL] Received null pointer from Go!" << std::endl;
    }
}

// 3. C++からGoへ文字列を受信する関数 (C++側で確保したメモリのポインタを返す方式)
const char* GetStringFromCPP() {
    std::string response = "Hello! This is a message from C++ DLL (Pointer method).";
    
    // Go側でFreeCPPStringを呼び出して解放してもらうため、ヒープに新しくメモリを確保する
    char* buffer = new char[response.length() + 1];
    std::memcpy(buffer, response.c_str(), response.length());
    buffer[response.length()] = '\0';
    
    return buffer;
}


