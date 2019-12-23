#include <iostream>
#include <list>
using namespace std;
struct lmnt
{
    char op;
    long long dgt;
    int un;
};
int N, opn; // 입력 문자열의 길이
list<lmnt> org, cpy, prt;
long long pa, a, ma = -(1 << 31);
// 일부 항의 연산 결과 partial answer, 모든 항의 연산 결과 answer, a들 중 최댓값 maximum answer
void calc(long long &a, char op, long long b)
{ // a = a (op) b
    switch (op)
    {
    case '+':
        a += b;
        break;
    case '-':
        a -= b;
        break;
    case '*':
        a *= b;
        break;
    }
}
void calc_all()
{
    cpy = org;
    for (int un = opn; un >= -opn; --un)
    {
        for (auto cpy_it = cpy.begin(); cpy_it != cpy.end();)
        {
            if ((*cpy_it).un != un)
            { // 현재 괄호 번호에 해당하지 않으면 continue
                ++cpy_it;
                continue;
            }
            auto prt_beg = cpy_it;                                      // prt의 시작 위치 기록
            prt.clear();                                                // prt 비우기
            for (; cpy_it != cpy.end() && (*cpy_it).un == un; ++cpy_it) // prt 채우기 : 현재 괄호 번호에 해당하는 항들
                prt.push_back(*cpy_it);
            int prt_size = prt.size(); // prt 크기 기록
            // * 계산
            for (auto prt_it = prt.begin(); prt_it != prt.end();)
            {
                auto prt_nit = prt_it;
                if (++prt_nit == prt.end())
                    break;
                if ((*prt_nit).op == '*')
                {
                    (*prt_it).dgt = (*prt_it).dgt * (*prt_nit).dgt;
                    prt.erase(prt_nit);
                }
                else
                    ++prt_it;
            } // +, - 계산
            auto prt_it = prt.begin();
            a = (*prt_it).dgt;
            for (++prt_it; prt_it != prt.end(); ++prt_it)
            {
                calc(a, (*prt_it).op, (*prt_it).dgt);
            } // 항들 계산후 결과값을 시작항에 할당, 나머지 항들은 erase
            (*prt_beg).dgt = a;
            --(*prt_beg).un;
            auto prt_begn = prt_beg;
            ++prt_begn;
            while (--prt_size)
                prt_begn = cpy.erase(prt_begn);
        }
    }
    if (a > ma)
        ma = a;
}
void BF(list<lmnt>::iterator it, int un)
{
    if (it == org.end())
    {
        calc_all();
        return;
    }
    auto nit = it;
    ++nit;
    (*it).un = un - 1;
    BF(nit, un - 1);
    (*it).un = un;
    BF(nit, un);
    (*it).un = un + 1;
    BF(nit, un + 1);
}
int main()
{
    ios_base::sync_with_stdio(0);
    cin.tie(0);
    cin >> N;
    cin.get();
    org.push_back({'\0', cin.get() - '0', 0});
    opn = (N + 1) / 2 - 1;
    for (int i = 0; i < opn; ++i)
        org.push_back({(char)cin.get(), cin.get() - '0', 0});
    auto it = org.begin();
    BF(++it, 0);
    cout << ma;
    return 0;
}
