# https://www.acmicpc.net/problem/17070

import fileinput
import time

global N
global m

class pipe:
    def __init__(self, x, y, stat): # x, y: coordination of head, stat: 'LONG', 'LATI', 'DIAG'
        self.x = x
        self.y = y
        self.stat = stat
    def __str__(self):
        return '({}, {})'.format(self.x, self.y)

# def x_move(p:pipe) -> pipe:
#     global N
#     global m
#     if p.stat == 'LATI':
#         return None
#     elif p.x + 1 >= N:
#         return None
#     elif m[p.y][p.x+1] != 0:
#         return None
#     else:
#         return pipe(p.x+1, p.y, 'LONG')

# def y_move(p:pipe) -> pipe:
#     global N
#     global m
#     if p.stat == 'LONG':
#         return None
#     elif p.y + 1 >= N:
#         return None
#     elif m[p.y+1][p.x] != 0:
#         return None
#     else:
#         return pipe(p.x, p.y+1, 'LATI')
    
# def xy_move(p:pipe) -> pipe:
#     global N
#     global m
#     if p.x + 1 >= N or p.y + 1 >= N:
#         return None
#     elif m[p.y+1][p.x+1] != 0 or m[p.y+1][p.x+1] != 0 or m[p.y+1][p.x+1] != 0:
#         return None
#     else:
#         return pipe(p.x+1, p.y+1, 'DIAG')

# def search(p:pipe) -> int:
#     if p:
#         if p.x == N-1 and p.y == N-1:
#             return 1
#         else:
#             if p.stat == "DIAG":
#                 return search(x_move(p)) + search(y_move(p)) + search(xy_move(p))
#             elif p.stat == "LONG":
#                 return search(x_move(p)) + search(xy_move(p))
#             else:
#                 return search(y_move(p)) + search(xy_move(p))
#     else:
#         return 0

def search(p:pipe) ->int:
    global N
    global m
    x=0
    y=0
    xy=0
    if p.stat == 'LONG':
        if p.x + 1 >= N:
            x = 0
        elif m[p.y][p.x+1] == 0:
            x = search(pipe(p.x+1, p.y, 'LONG'))
        else:
            x = 0
        if p.x + 1 >= N or p.y + 1 >= N:
            xy = 0
        elif m[p.y][p.x+1] == 0 and m[p.y+1][p.x] == 0 and m[p.y+1][p.x+1] == 0:
            xy = search(pipe(p.x+1, p.y+1, 'DIAG'))
        else:
            xy = 0
    elif p.stat == 'LATI':
        if p.y + 1 >= N-2:
            y= 0
        elif m[p.y+1][p.x] == 0:
            y = search(pipe(p.x, p.y+1, 'LATI'))
        else:
            y = 0
        if p.x + 1 >= N or p.y + 1 >= N:
            xy = 0
        elif m[p.y][p.x+1] == 0 and m[p.y+1][p.x] == 0 and m[p.y+1][p.x+1] == 0:
            xy = search(pipe(p.x+1, p.y+1, 'DIAG'))
        else:
            xy = 0
    elif p.stat == 'DIAG':
        if p.x + 1 >= N:
            x = 0
        elif m[p.y][p.x+1] == 0:
            x = search(pipe(p.x+1, p.y, 'LONG'))
        else:
            x = 0
        if p.y + 1 >= N:
            y= 0
        elif m[p.y+1][p.x] == 0:
            y = search(pipe(p.x, p.y+1, 'LATI'))
        else:
            y = 0
        if p.x + 1 >= N or p.y + 1 >= N:
            xy = 0
        elif m[p.y][p.x+1] == 0 and m[p.y+1][p.x] == 0 and m[p.y+1][p.x+1] == 0:
            xy = search(pipe(p.x+1, p.y+1, 'DIAG'))
        else:
            xy = 0
    return x+y+xy


def main():
    global N, m
    stdin = fileinput.input()
    N = int(stdin.readline().strip())
    m = []
    for ___ in range(N):
        __ = []
        _ = stdin.readline().split(' ')
        for c in _:
            __.append(int(c))
        m.append(__)
    start = time.time()
    first = pipe(0,1,'LATI')
    print(search(first))
    end = time.time()
    print(end-start)
        
        

if __name__ == '__main__':
    main()