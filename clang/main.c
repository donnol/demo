#include <stdio.h>

extern int player(); // 不写这一行的话，在C99标准编译时会有警告：warning: implicit declaration of function ‘player’ [-Wimplicit-function-declaration]

int fun(void){
    // static 在局部变量
    // 这条赋值语句虽然写在这里，但其实没有在代码编译后的机器指令里存在。
    // 既然赋值语句都没有执行，它的初始值为什么会是10呢？
    // 事实上在编译阶段，就把可执行文件中为这个变量分配的静态存储区的值给赋成了10，程序在加载的时候将10直接读入内存，而不会有“赋值”这个动作了。
    static int count = 10;    // 事实上此赋值语句从来没有执行过
    return count--;
}

int count = 1;

int main(void)
{
    // printf("hello world.\n");

    // if ('abc' < 'ABC')
    // {
    //     printf("true\n");
    // }
    // else
    // {
    //     printf("false\n");
    // }

    // // if ((3 is 4) == 0)
    // // {
    // //     printf("true\n");
    // // }
    // // else
    // // {
    // //     printf("false\n");
    // // }

    // if (8 > 4 > 2)
    // {
    //     printf("true\n");
    // }
    // else
    // {
    //     printf("false\n");
    // }

    // // if (9 < 1 and 10 > 9 or 2 > 1)
    // // {
    // //     printf("true\n");
    // // }
    // // else
    // // {
    // //     printf("false\n");
    // // }

    // // enum
    // enum Color{RED, BLUE, GREEN}; // 声明 Color 枚举类型
    // printf("%d, %d, %d\n", RED, BLUE, GREEN); // RED, BLUE, GREEN 都是整型常量
    // enum Color red=RED, blue=BLUE, green=GREEN; // 声明枚举类型变量并初始化
    // enum Color black; // 只声明不初始化，默认是 0，为什么有时会输出 4195472？
    // printf("%d, %d, %d, %d\n", red, blue, green, black);
    // int r  = red; // 赋值给整型变量
    // printf("%d\n", r);

    // enum Weekday{Monday="monday"}; // 必须是整型常量才行
    // enum Weekday monday = Monday;
    // printf("%s\n", monday);

    // union
    union Union{int a; float f; char *c;};
    union Union u;

    // 同一时间只能储存一个数据成员
    // u.a = 1;
    // printf("%d\n", u.a);
    // u.f = 1.0;
    // printf("%f\n", u.f);
    u.c = "hello";
    printf("%s\n", u.c);

    // 分配最大长度，这里是 8
    // 8, 4, 8, 4
    printf("%ld, %ld, %ld, %ld\n", sizeof(u), sizeof(u.a), sizeof(u.c), sizeof(u.f));

    // extern
    extern int game; // 在其它文件声明的全局变量
    printf("%d\n", game);

    // register
    register int reg = 10; // 保存在 CPU 寄存器
    printf("%d\n", reg);

    // static
    // 隐藏
    int pl = player();
    printf("%d\n", pl);

    // 持久性
    printf("global\t\tlocal static\n");
    for(; count <= 10; ++count)
        printf("%d\t\t%d\n", count, fun());    
    
    // 默认值0
    static char str[10];
    printf("'%s'\n", str);

    // restrict 需要在编译时添加标识 -std=c99 开启 C99 支持
    int *restrict res;
    res = &pl;
    printf("%d\n", *res);

    return 0;
}