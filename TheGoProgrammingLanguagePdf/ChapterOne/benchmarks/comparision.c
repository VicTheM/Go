#include <stdio.h>

int main(void) {
    for (int x = 0; x < 1000000; x++) {
        printf("routine iteration: %d", x);
    }

    return 0;
}