
g++ -shared -I./include -I/prog/vcpkg/installed/x64-windows/include ^
 -L/prog/vcpkg/installed/x64-windows/lib ^
 -o mydll.dll dll.cpp -Wl,--out-implib,libmydll.a

