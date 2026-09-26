# Codeforces 1124A - SauSaGe Bank

- **Problem Link**: [1124A - SauSaGe Bank](https://codeforces.com/contest/1124/problem/A)
- **Submission ID**: [#392180993](https://codeforces.com/contest/1124/submission/392180993)
- **Language**: C++
- **Verdict**: Accepted

## Solution Code

```cpp
t = int(input()) for i in range(t):  n , k = map(int,input().split())  total = 0  total += (k - 1) * 2  rem_days = n - (k - 1)  total += 2 ** (rem_days)  print(total) 
```
