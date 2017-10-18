#include <stdio.h>
/***
	Function performs the product of numbers 0 and n.
***/
int prod (int n){
	if(n==0){
		return 0;
	}else if (n==1){
		return 1;
	}else{
		return n*sum(n-1);
	}
}
int main(){
	int n;
	scanf("%d", &n);
	printf("Value: %d \nProduct: %d\n",n,prod(n));
	return 0;
}