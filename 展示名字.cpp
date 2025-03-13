/*展示名字*/
/*
皮皮的班上有五个同学，
开始的时候分别录入每个同学的姓，
依次展示这些同学的姓。
后来发现有一些姓是重复的，
就想把名补充上，
展示完整的姓名，
请你写个程序帮皮皮来完成任务吧！
*/
#include <iostream>
#include <string>
using namespace std;

int main()
{
	string s1[10];
	string s2[10];
	for (int i = 0;i < 5;i++)
	{
		cin>>s1[i];
	}
	for (int i = 0;i < 5;i++)
	{
		cin>>s2[i];
		cout<<s1[i] + s2[i]<<" ";
	}
	return 0;
}