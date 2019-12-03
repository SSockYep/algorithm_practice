# https://www.acmicpc.net/problem/17135
import itertools

global n, m, d, castle, count

def shoot(turn, pos):
    global castle, count, n, m, d
    for i in range(0, d):
        y = n-turn
        x = pos-i
        while x <= min(pos+i, m-1) and y <= n-turn:
            if y < 0 or x < 0:
                pass
            elif castle[y][x] == 1:
                if not [y, x] in count:
                    count.append([y, x])
                return
            if x < pos:
                y -= 1
            else:
                y += 1
            x += 1
            
def main():
    global n, m, d, castle, count
    n, m, d = input().split(' ')
    n = int(n)
    m = int(m)
    d = int(d)
    castle = []
    count = []
    res = 0
    for i in range(n):
        _ = list(map(int, input().split(' ')))
        castle.append(_)

    for c1, c2, c3 in itertools.combinations(range(m), 3):
        for i in range(1, n+1):
            shoot(i,c1)
            shoot(i,c2)
            shoot(i,c3)
            for ss in count:
                castle[ss[0]][ss[1]] = 0
                
        res = max(res, len(count))
        for ss in count:
            castle[ss[0]][ss[1]] = 1
        count = []
    print(res)

if __name__ == '__main__':
    main()