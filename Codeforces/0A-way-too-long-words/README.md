# Codeforces 0A - Way Too Long Words

- **Problem Link**: [0A - Way Too Long Words](https://codeforces.com/contest/0/problem/A)
- **Submission ID**: [#376807712](https://codeforces.com/submissions/Greninja22)
- **Language**: go
- **Verdict**: Accepted

## Solution Code

```go
t = int(input())
 
for i in range(t):
  s = input()
  first = s[0]
  last = s[-1]
  if len(s) > 10:
    print(first + str(len(s) - 2) + last)
  else:
    print(s)
  
```
