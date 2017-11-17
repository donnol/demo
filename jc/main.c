#include <stdio.h>

int main(void)
{
    printf("hello world.\n");

    if ('abc' < 'ABC')
    {
        printf("true\n");
    }
    else
    {
        printf("false\n");
    }

    // if ((3 is 4) == 0)
    // {
    //     printf("true\n");
    // }
    // else
    // {
    //     printf("false\n");
    // }

    if (8 > 4 > 2)
    {
        printf("true\n");
    }
    else
    {
        printf("false\n");
    }

    // if (9 < 1 and 10 > 9 or 2 > 1)
    // {
    //     printf("true\n");
    // }
    // else
    // {
    //     printf("false\n");
    // }

    return 0;
}