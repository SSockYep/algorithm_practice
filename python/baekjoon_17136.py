# https://www.acmicpc.net/problem/17136

global r, rest

def search(x, y, size, d, m):
    global r, rest
    _m = []
    for l in m:
        _m.append(l[:])
    for l in _m:
        if '1' in l:
            break
    else:
        r = min(r, d-1)
        return

    if y+size > 10 or x+size > 10 or x >= 10 or y >= 10 or d-1 >= r:
        return
    else:
        flag = False
        for l in _m[y:y+size]:
            if '0' in l[x:x+size]:
                flag = True
                break
        if flag:
            return
        else:
            for i in range(y, y+size):
                _m[i][x:x+size] = ['0'] * size

    y2=0
    flag=False
    while y2 < 10:
        x2 = 0
        while x2 < 10:
            if _m[y2][x2] == '1':
                flag=True
                break
            x2 += 1
        if flag:    
            break
        y2 += 1

    for i in range(5,0,-1):
        if rest[i-1] == 5:
            continue
        rest[i-1] += 1
        search(x2, y2, i, d+1, _m)  
        rest[i-1] -= 1

def main():
    global r, rest
    m = []
    rest = [0,0,0,0,0]
    r = 26
    for _ in range(10):
        m.append(input().split(' '))

    search(0,0,0,0,m)
    
    if r == 26:
        print(-1)
    elif r == -1:
        print(0)
    else:
        print(r)
    

if __name__ == '__main__':
    main()