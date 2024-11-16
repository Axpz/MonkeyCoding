/*
 题目描述

皮皮在摘香蕉的时候，没能忍住，有时会偷懒躲在家里不去摘香蕉，反而还吃掉若干香蕉。每天香蕉的变化数量为xi xi ，xi xi为正数时，表示皮皮去摘了xixi根香蕉，否则表示皮皮躲在家里偷懒，并吃了∣xi∣∣xi ∣ 根香蕉。
求当香蕉剩余数量第一次不少于 
1000
1000 的时候，皮皮一共摘了多少香蕉。
*/

#include <iostream>
using namespace std;

int main()
{
    int n = 0, m = 0, nn;
    while (cin >> nn)
    {
        if (nn > 0)
        {
            m += nn;
        }
        n += nn;
        if (n >= 1000)
        {
            break;
        }
    }
    cout << m;
    return 0;
}