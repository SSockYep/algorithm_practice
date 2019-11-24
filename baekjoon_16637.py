# https://www.acmicpc.net/problem/16637

import fileinput

global N
global ans
global exp
global parantheses

def calc(lhs, sign, rhs) -> int:
    if sign == '+':
        return lhs + rhs
    elif sign == '-':
        return lhs - rhs
    elif sign == '*':
        return lhs * rhs

def cal_exp(exp) -> int:
    st = []
    i = 0
    global parantheses
    while i < len(exp):
        if parantheses[i] == 0:
            st.append(exp[i])
            i += 1
        else:
            st.append(calc(int(exp[i]), exp[i+1], int(exp[i+2])))
            i = i + 3
    st.reverse()
    while len(st) > 1:
        l = int(st.pop())
        sign = st.pop()
        r = int(st.pop())
        st.append(calc(l, sign, r))
    return int(st[0])

def max(a,b):
    if a > b:
        return a
    else:
        return b

def dfs(index):
    global ans, N, exp, parantheses
    ans = max(ans, cal_exp(exp))
    for i in range(index, N+1, 2):
        if i+2 <= N:
            if parantheses[i] == 0 and parantheses[i+2] == 0:
                parantheses[i] = 1
                parantheses[i+2] = 1
                dfs(i+2)
                parantheses[i] = 0
                parantheses[i+2] = 0



def main():
    global N
    global ans
    global exp, parantheses
    stdin = fileinput.input()
    N = int(stdin.readline())
    exp = list(stdin.readline().strip())
    ans = - 2 ** 31
    parantheses = [0] * N
    
    dfs(0)
    
        
    print(ans)


if __name__ == '__main__':
    main()
