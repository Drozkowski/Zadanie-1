#include <iostream>

using namespace std;

int main() {
    int a = 1;
    int b = 0;

    while (a != 0) {
        cout << "Wprowadz liczbe: ";
        cin >> a;

        if (b > a) {
            b = a;
        }
    }

    cout << "Najmniejsza liczba: " << b << endl;

    return 0;
}