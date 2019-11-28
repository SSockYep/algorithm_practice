# https://www.acmicpc.net/problem/17070

import fileinput

global N
global m
global result

def search(x:int, y:int, d:int):
    global result

    if x == N-1 and y == N-1:
        result = result + 1
    else:
        if y + 1 < N and x + 1 < N:
            if m[y+1][x+1] == 0 and m[y][x+1] == 0 and m[y+1][x] == 0:
                search(x+1, y+1, 2)
        if x + 1 < N:
            if (d == 0 or d == 2) and m[y][x+1] == 0:
                search(x+1, y, 0)
        if y + 1 < N:
            if (d == 1 or d == 2) and m[y+1][x] == 0:
                search(x, y+1, 1)

def main():
    global N, m
    global result
    result=0
    stdin = fileinput.input()
    N = int(stdin.readline().strip())
    m = []
    for ___ in range(N):
        __ = []
        _ = stdin.readline().split(' ')
        for c in _:
            __.append(int(c))
        m.append(__)
    search(1,0,0)
    print(result)
        
        

if __name__ == '__main__':
    main()