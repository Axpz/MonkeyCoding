/*
“百鸡问题”是出自我国古代《张丘建算经》的著名数学问题。大意为:“每只公鸡 5 元，每只母鸡 3 元，每 3 只小鸡 1 元; 现在有 100 元，买了 100 只鸡，共有多少种方案?”
小明很喜欢这个故事，他决定对这个问题进行扩展，并使用编程解决:如果 每只公鸡x元，每只母鸡y元，每z只小鸡1元;现在有n元，买了m只鸡，共有多少种方案?
*/ 
#include <iostream>
using namespace std;

int main()
{
    int cnt = 0;
    int x,y,z,n,m;
    cin>>x>>y>>z>>n>>m;
    for(int i = 0;i <= m;i++)
    {
        for(int j = 0;j <= m - i;j++)
        {
            int k = m - i - j;
            if(k % z == 0)
            {
                if(x * i + y * j + k / z == n)
                {
                    cnt++;
                }
            }
        }
    }
    cout<<cnt;
}