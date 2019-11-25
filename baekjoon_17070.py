# https://www.acmicpc.net/problem/17070

import fileinput

global N
global m

class pipe:
    def __init__(self, x, y, stat): # x, y: coordination of head, stat: 'LONG', 'LATI', 'DIAG'
        self.x = x
        self.y = y
        self.stat = stat
    def __str__(self):
        return '({}, {})'.format(self.x, self.y)

def x_move(p:pipe) -> pipe:
    global N
    global m
    if p.stat == 'LATI':
        return None
    elif p.x + 1 >= N:
        return None
    elif m[p.y][p.x+1] != 0:
        return None
    else:
        return pipe(p.x+1, p.y, 'LONG')

def y_move(p:pipe) -> pipe:
    global N
    global m
    if p.stat == 'LONG':
        return None
    elif p.y + 1 >= N:
        return None
    elif m[p.y+1][p.x] != 0:
        return None
    else:
        return pipe(p.x, p.y+1, 'LATI')
    
def xy_move(p:pipe) -> pipe:
    global N
    global m
    if p.x + 1 >= N or p.y + 1 >= N:
        return None
    elif m[p.y+1][p.x+1] != 0 or m[p.y+1][p.x+1] != 0 or m[p.y+1][p.x+1] != 0:
        return None
    else:
        return pipe(p.x+1, p.y+1, 'DIAG')

def search(p:pipe) -> int:
    if p:
        if p.x == N-1 and p.y == N-1:
            return 1
        else:
            return search(x_move(p))+search(y_move(p))+search(xy_move(p))
    else:
        return 0

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
    first = pipe(0,1,'LATI')
    print(search(first))
        
        

if __name__ == '__main__':
    main()