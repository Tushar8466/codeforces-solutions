# Codeforces 2269A - SauSaGe Bank

- **Problem Link**: [2269A - SauSaGe Bank](https://codeforces.com/contest/2269/problem/A)
- **Submission ID**: [#1790444262226](https://codeforces.com/contest/2269/submission/1790444262226)
- **Language**: C++
- **Verdict**: Accepted

## Solution Code

```cpp
t = int(input()) for i in range(t):  n , k = map(int,input().split())  total = 0  total += (k - 1) * 2  rem_days = n - (k - 1)  total += 2 ** (rem_days)  print(total) 
```
