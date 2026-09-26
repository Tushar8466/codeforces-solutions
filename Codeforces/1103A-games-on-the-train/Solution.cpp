#include <iostream>
using namespace std;
int main() {
    int t;
    cin >> t;
    while (t--) {
        int n;
        cin >> n;
        int mx = 0, mn = 100;
        for (int i = 0; i < n; i++) {
            int arr;
            cin >> arr;
            mx = max(mx, arr);
            mn = min(mn, arr);
        }
        cout << (mx - mn + 1) << "\n";
    }
}