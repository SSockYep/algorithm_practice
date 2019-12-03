// https://www.acmicpc.net/problem/17136

#include <iostream>
#include <vector>

using namespace std;

int r = 26;
int rest[5] = {0,0,0,0,0};

const inline int min(int a, int b){
    return a < b ? a : b;
}

void search(int x, int y, int size, int d, vector<vector <int>> m) {

    bool check = false;
    for (int i = 0; i < 10; i++){
        for (int j = 0; j < 10; j++){
            if (m[i][j] == 1){
                check = true;
                break;
            }
        }
        if (check){
            break;
        }
    }
    
    if (!check) {
        r = min(r, d-1);
        return;
    }

    if (y+size > 10 || x+size > 10 || d-1 >= r)
        return;
    else {
        bool flag = false;
        for(int i = y; i < y+size; i++) {
            for (int j = x; j < x+size; j++) {
                if (m[i][j]==0){
                    return;
                }
            }
        }
        for(int i = y; i < y+size; i++) {
            for (int j = x; j < x+size; j++) {
                m[i][j] = 0;
            }
        }
    }

    int x2, y2;
    bool flag=false;
    for(y2 =0; y2<10; y2++){
        for (x2 = 0; x2<10; x2++){
            if (m[y2][x2]==1){
                flag = true;
                break;
            }
        }
        if(flag) {
            break;
        }
    }

    for (int i = 5; i > 0; i--) {
        if (rest[i-1] == 5) continue;
        rest[i-1]++;
        search(x2,y2,i,d+1,m);
        rest[i-1]--;
    }
}

int main() {
    vector<vector <int>> m;
    for(int _=0; _<10; _++){
        vector <int> tmp;
        for (int __=0; __<10; __++) {
            int c;
            cin >> c;
            tmp.push_back(c);
        }
        m.push_back(tmp);
    }

    search(0,0,0,0,m);
    
    if (r == 26) {
        cout << -1 << endl;
    }
    else if (r == -1) {
        cout << 0 << endl;

    }
    else {
        cout << r << endl;
    }
}
