#include <iostream>
#include <algorithm>
using namespace std;
int main(){
  int t;
  cin >> t;
  while(t--){
    int n;
    cin >> n;
 
    int a[100];
    for(int i = 0; i < n; i ++){
      cin >> a[i];
    }
    sort(a , a + n);
    int mid = a[n / 2];
    int left = 0;
    int right = 0;
    for(int i = 0; i < n ; i++){
      if(a[i] < mid){
        left ++;
      }
      else if(a[i] > mid){
        right ++;
      }
    }
    cout << max(left , right) << endl;
  }
  return 0;
}