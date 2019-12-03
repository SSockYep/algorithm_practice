// # https://www.acmicpc.net/problem/17281

#include <iostream>
#include <algorithm>
#include <vector>

using namespace std;
// from itertools import permutations
// import time


// def main():
int main() {
    int n;
    vector<vector <int>> players;
    int ans = 0;
    cin >> n;
    players = vector<vector <int>>(n, vector<int>(9));

    // n = int(input())
    // players = []
    // ans = 0
    
    for (int i = 0; i < n; i++) {
        for (int j = 0; j < 9; j++) {
            cin >> players[i][j];
        }
    }
    // for _ in range(n):
    //     tmp = list(map(int, input().strip().split(' ')))
    //     players.append(tmp)

    // start=time.time()
    vector<int> perm = {1,2,3,4,5,6,7,8};
    do {
        // for perm in permutations(range(1,9)):
        vector<int>base = {0,0,0};
        int inning = 0, i = 0, ret = 0;
        // base = [0,0,0]
        // inning = 0
        // i = 0
        // ret = 0
        while(inning < n) {
            // while inning < n:
            int outcount = 0;
            base = {0,0,0};
            int score = 0;
            // outcount = 0
            // base = [0,0,0]
            // score = 0
            while (outcount < 3){
            // while outcount < 3:
            int play;
                if(i == 3) {
                    play = players[inning][0];
                }
                else if (i < 3)
                {
                    play = players[inning][perm[i]];
                }
                else {
                    play = players[inning][perm[i-1]];
                }
                
                // if i == 3:
                //     play = players[inning][0]
                // elif i < 3:
                //     play = players[inning][perm[i]]
                // else:
                //     play = players[inning][perm[i-1]]
                if (play == 0) {
                    outcount++;
                }
                else if (play == 1) {
                    if(base[2]){
                        score++;
                    }
                    base[2] = base[1];
                    base[1] = base[0];
                    base[0] = 1;
                }
                else if (play == 2) {
                    if(base[2]){
                        score++;
                    }
                    if(base[1]){
                        score++;
                    }
                    base[2] = base[0];
                    base[1] = 1;
                    base[0] = 0;
                }
                else if (play == 3) {
                    if(base[2]){
                        score++;
                    }
                    if(base[1]){
                        score++;
                    }
                    if(base[0]){
                        score++;
                    }
                    base[2] = 1;
                    base[1] = 0;
                    base[0] = 0;
                }
                else if (play == 4) {
                    if(base[2]){
                        score++;
                    }
                    if(base[1]){
                        score++;
                    }
                    if(base[0]){
                        score++;
                    }
                    score++;
                    base = {0,0,0};
                }
                // if play == 0:
                //     outcount += 1
                // elif play == 1:
                //     if base[2]:
                //         score += 1
                //     base[2] = base[1]
                //     base[1] = base[0]
                //     base[0] = 1
                // elif play == 2:
                //     if base[2]:
                //         score += 1
                //     if base[1]:
                //         score += 1
                //     base[2] = base[0]
                //     base[1] = 1
                //     base[0] = 0
                // elif play == 3:
                //     if base[2]:
                //         score += 1
                //     if base[1]:
                //         score += 1
                //     if base[0]:
                //         score += 1
                //     base[2] = 1
                //     base[1] = 0
                //     base[0] = 0
                // elif play == 4:
                //     if base[2]:
                //         score += 1
                //     if base[1]:
                //         score += 1
                //     if base[0]:
                //         score += 1
                    // score += 1
                    // base = [0,0,0]
                i++;
                if (i == 9){
                    i = 0;
                }
                // i += 1
                // if i == 9:
                //     i = 0
            }
            ret += score;
            inning++;
            // ret += score
            // inning += 1
        }
        ans = ans > ret ? ans : ret;
    } while(next_permutation(perm.begin(), perm.end()));

    cout << ans << endl;

    return 0;
    
    // print(ans)
    // print(time.time() - start)

}
// if __name__ == '__main__':
//     main()