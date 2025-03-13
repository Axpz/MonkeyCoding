/*
小明刚刚学习了素数的概念：如果一个大于 1 的正整数，除了1 和它自身外，不能被其他正整数整除，则这个正整数是素数。
现在，小明想找到两个正整数A和B之间（包括A和B）有多少个素数？
*/
#include <iostream>
using namespace std;

bool sushu(int a)
{
	for(int i = 2;i <= a - 1;i++)
	{
		if(a % i == 0)
		{
			return false;
		}
	}
	return true;
}

int main()
{
	int a,b;
	cin>>a>>b;
	int cnt = 0;
	for(int i = a;i <= b;i++)
	{
		if(sushu(i))
		{
			cnt++;
		}
	}
	cout<<cnt;
}