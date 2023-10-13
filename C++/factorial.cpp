//  Created by Riya Pannu on 15/10/17.

#include<iostream>
using namespace std;

// factorial fucntion declaration
long long int factorial(long long int n) {
    if (n<0) //User input less than -1 error handling
        return -1;
    else if (n == 0) //Break Condition for recursion 
        return 1;
    else
        return n * factorial(n-1); //Calling factorial function inside factorial function : f(f(x)), is example of recursions.
    
}

int main() {
    long long int n; // Decleared long dataType 'n' variable
    cout<<"Enter a positive number \n";
    cin >> n; // Taking Input from User (0-25)
    long long int result = factorial(n);
    if (result==-1)
        cout<<"Invalid input";
    cout << "The factorial of "<<n<<" is : "<<result << endl;//Print Statement 
    return 0;
}
