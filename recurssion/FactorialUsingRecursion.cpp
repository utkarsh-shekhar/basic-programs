#include <iostream>
using namespace std;

float fact(int n)
{
	int res;
	if(n>1)
		return n*fact(n-1);
	else
		return 1;
}

int main()
{
	int t;
	cin>>t;
	while(t--)
	{
		int n;
		cin>>n;
		cout<<fact(n)<<endl;
	}
	return 0;
}
