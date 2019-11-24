# https://www.acmicpc.net/problem/16637

import fileinput

def calc(lhs, sign, rhs) -> int:
    if sign == '+':
        return lhs + rhs
    elif sign == '-':
        return lhs - rhs
    elif sign == '*':
        return lhs * rhs

def main():
    N=0
    exp = []
    numbers = []
    signs = []
    ans = 0
    stdin = fileinput.input()
    N= int(stdin.readline())
    exp = list(stdin.readline().strip())
    if N < 1 or N > 19 or N % 2 == 0:
        return -1
    elif len(exp) != N:
        return -1

    for i in range(0, N):
        if i == 0:
            ans = int(exp[i])
        elif i % 2 == 0:
            numbers.append(int(exp[i]))
        else:
            signs.append(exp[i])
    # print(numbers)
    # print(signs)
    # print(ans)
    i = 0
    while i < (N-1)/2 - 1:
        if signs[i+1] == '-':
            if i+2 < (N-1)/2:
                if signs[i+2] == '-':
                    first =  calc(ans, signs[i], numbers[i]) - (numbers[i+1] - numbers[i+2])
                    second = calc(ans, signs[i], (numbers[i] - numbers[i+1])) - numbers[i+2]
                    if first > second:
                        ans = first
                        # print('%s%d%s(%d%s%d)=%d'%(signs[i],numbers[i],signs[i+1],numbers[i+1],signs[i+2],numbers[i+2],ans))
                        i += 3
                    else:
                        ans = calc(ans, signs[i], (numbers[i] - numbers[i+1]))
                        # print('%s(%d%s%d)=%d'%(signs[i],numbers[i],signs[i+1],numbers[i+1],ans))
                        i += 2
                else:
                    np = calc(calc(ans, signs[i], numbers[i]), signs[i+1], numbers[i+1])
                    p = calc(ans, signs[i], calc(numbers[i], signs[i+1], numbers[i+1]))
                    if  p < np:
                        ans = calc(ans, signs[i], numbers[i])
                        # print('%s%d=%d'%(signs[i],numbers[i],ans))
                        i += 1
                    else:
                        ans = p
                        # print('%s(%d%s%d)=%d'%(signs[i],numbers[i],signs[i+1],numbers[i+1],ans))
                        i += 2
            else: 
                np = calc(calc(ans, signs[i], numbers[i]), signs[i+1], numbers[i+1])
                p = calc(ans, signs[i], calc(numbers[i], signs[i+1], numbers[i+1]))
                if  p < np:
                    ans = calc(ans, signs[i], numbers[i])
                    # print('%s%d=%d'%(signs[i],numbers[i],ans))
                    i += 1
                else:
                    ans = p
                    # print('%s(%d%s%d)=%d'%(signs[i],numbers[i],signs[i+1],numbers[i+1],ans))
                    i += 2
                    
        else:
            np = calc(calc(ans, signs[i], numbers[i]), signs[i+1], numbers[i+1])
            p = calc(ans, signs[i], calc(numbers[i], signs[i+1], numbers[i+1]))
            if  p < np:
                ans = calc(ans, signs[i], numbers[i])
                # print('%s%d=%d'%(signs[i],numbers[i],ans))
                i += 1
            else:
                ans = p
                # print('%s(%d%s%d)=%d'%(signs[i],numbers[i],signs[i+1],numbers[i+1],ans))
                i += 2
        # print(ans)


    while i < int((N-1)/2):
        ans = calc(ans, signs[i], numbers[i])
        i += 1
    print(ans)


if __name__ == '__main__':
    main()
