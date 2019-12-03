# https://www.acmicpc.net/problem/17281

from itertools import permutations
import time

def calc(n:int, perm:tuple, players:list) -> int:
    
    return ret

def main():
    n = int(input())
    players = []
    ans = 0

    for _ in range(n):
        tmp = list(map(int, input().strip().split(' ')))
        players.append(tmp)

    start=time.time()

    for perm in permutations(range(1,9)):
        base = [0,0,0]
        inning = 0
        i = 0
        ret = 0
        while inning < n:
            outcount = 0
            base = [0,0,0]
            score = 0
            while outcount < 3:
                if i == 3:
                    play = players[inning][0]
                elif i < 3:
                    play = players[inning][perm[i]]
                else:
                    play = players[inning][perm[i-1]]

                if play == 0:
                    outcount += 1
                elif play == 1:
                    if base[2]:
                        score += 1
                    base[2] = base[1]
                    base[1] = base[0]
                    base[0] = 1
                elif play == 2:
                    if base[2]:
                        score += 1
                    if base[1]:
                        score += 1
                    base[2] = base[0]
                    base[1] = 1
                    base[0] = 0
                elif play == 3:
                    if base[2]:
                        score += 1
                    if base[1]:
                        score += 1
                    if base[0]:
                        score += 1
                    base[2] = 1
                    base[1] = 0
                    base[0] = 0
                elif play == 4:
                    if base[2]:
                        score += 1
                    if base[1]:
                        score += 1
                    if base[0]:
                        score += 1
                    score += 1
                    base = [0,0,0]

                i += 1
                if i == 9:
                    i = 0
            ret += score
            inning += 1
        ans = max(ans, ret)
    
    print(ans)
    print(time.time() - start)

if __name__ == '__main__':
    main()