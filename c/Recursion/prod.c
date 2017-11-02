#include <stdio.h>
/***
	Function prod performs the product of numbers 1 and n or zero.
	Receives the last n (number) of the product.
	Returns the product from 1 to n.
***/
int prod (int n){
	if(n==0){
		return 0;
	}else if (n==1){
		return 1;
	}else{
		return n*prod(n-1);
	}
}


int main(){
	int n;
	scanf("%d", &n);
	
	if(n<0)
		printf("The program only performs the output for positive numbers!\n");
	else
		printf("Value: %d \nProduct: %d\n",n,prod(n));

	return 0;
}