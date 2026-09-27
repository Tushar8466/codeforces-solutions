t = int(input())
 
for i in range(t):
  s = input()
  first = s[0]
  last = s[-1]
  if len(s) > 10:
    print(first + str(len(s) - 2) + last)
  else:
    print(s)
  