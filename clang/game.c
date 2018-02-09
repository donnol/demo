// 当使用 static 声明全局变量/函数时，可以达到隐藏该变量/函数的作用，其它源文件无法使用它们
// static int game = 1; // 当全局变量用 static 声明后，该函数只能在本源文件内使用

// static int player() { // 当全局函数用 static 声明后，该函数只能在本源文件内使用
//     return game;
// }

int game = 1;

int player() {
    return game;
}