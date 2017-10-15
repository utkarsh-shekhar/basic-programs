#include<iostream>
#include<string>
using namespace std;

int main(){
	for(int i = 1;i <= 100;i++){
		string output = "";
		output += i % 3 == 0 ? "Fizz" : "";
		output += i % 5 == 0 ? "Buzz" : "";
		if(output == "") output = to_string(i);
		cout << output << endl;
	}
	return 0;
}