#include <iostream>
using namespace std;

int main() 
{
    //m heads, n legs
    int m,n;
    cin>>m>>n;
    for (int i=0; i<=m; i++)
    {
        for (int j=0; j<=m-i; j++)
        {
            int k = m-i-j;
            if (i*2+j*4+k*6==n)
            {
                cout<<i<<" "<<j<<" "<<k<<endl;
            }
        }
    }
    return 0;
}