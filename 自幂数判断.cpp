/*
自幂数是指，一个 N 位数，满足各位数字 N 次方之和是本身。例如，153 是 3 位数，其每位数的 3 次方之和，1 
3
 +5 
3
 +3 
3
 =153，因此 153 是自幂数；1634 是 4 位数，其每位数的 4 次方之和，1 
4
 +6 
4
 +3 
4
 +4 
4
 =1634，因此 1634 是自幂数。
现在，输入若干个正整数，请判断它们是否是自幂数。
*/
#include <iostream>
using namespace std;

int qiumi(int a,int b)
{
    int h = a;
    for(int i = 1;i <= b - 1;i++)
    {
        h *= a;
    }
    return h;
}

int m[11];


int main()
{
    int n;
    cin>>n;
    for(int i = 1;i <= n;i++)
    {
        int a;
        cin>>a;
        int t = a;
        int he = 0;
        int cnt = 0;
        while(t)
        {
            cnt++;
            m[cnt] = t % 10;
            t /= 10;
        }
        for(int j = 1;j <= cnt;j++)
        {
            he += qiumi(m[j],cnt);
        }
        if(he == a)
        {
            cout<<"T";
        }
        else
        {
            cout<<"F";
        }
        cout<<endl;
    }
}