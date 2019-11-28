#include <stdio.h>

#define MAX 16

int N;
int m[MAX][MAX];
int result = 0;

void search(int x, int y, int d) {
    if (x == N-1 && y == N-1) {
        result++;
    }
    else {
        if (y + 1 < N && x + 1 < N) {
            if (m[y+1][x+1] == 0 && m[y][x+1] == 0 && m[y+1][x] == 0) {
                search(x+1, y+1, 2);
            }
        }
        if (x + 1 < N) {
            if ((d == 0 || d == 2) && m[y][x+1] == 0) {
                search(x+1, y, 0);
            }
        }
        if (y + 1 < N) {
            if ((d == 1 || d == 2) && m[y+1][x] == 0) {
                search(x, y+1, 1);
            }
        }
    }
}

int main() {
    int i, j, t;
    scanf("%d", &N);
    for (i = 0; i < N; i++){
        for (j = 0; j < N; j++){
            scanf("%d", &t);
            m[i][j] = t;
        }
    }
    search(1,0,0);
    printf("%d\n", result);
    return 0;
}