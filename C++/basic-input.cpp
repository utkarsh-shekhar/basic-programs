#include <iostream>
#include <string>
using namespace std;

int main() {
    string s;
    int n;
    float f;
    double d;
    char c;
    cout << "Enter some text:" << endl;
    cin >> s;
    cout << "Enter an integer:" << endl;
    cin >> n;
    cout << "Enter a floating point number - precise up to 6 digits after comma:" << endl;
    cin >> f;
    cout << "Enter a floating point nubmer - precise up to 15 digits after comma:" << endl;
    cin >> d;
    cout << "Enter a single character:" << endl;
    cin >> c;

    cout << "Your string value is " << s << endl;
    cout << "Your int value is " << n << endl;
    cout << "Your float value is " << f << endl;
    cout << "Your double value is " << d << endl;
    cout << "Your char value is " << c << endl;
    return 0;
}
