
clang++ -shared -std=c++17 -I./include -I/prog/vcpkg/installed/x64-windows/include ^
 -L/prog/vcpkg/installed/x64-windows/lib -L./lib ^
 -lsqlite3 ^
 -o mydll.dll dll.cpp

