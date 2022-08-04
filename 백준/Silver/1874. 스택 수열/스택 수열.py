import sys 
input = sys.stdin.readline
sys.setrecursionlimit(10**6)

n = int(input())
num = 0
stack = []
result = []
ans=True

for i in range(n):
    x = int(input())
    while num < x:
      num += 1
      stack.append(num)
      result.append("+")

    if stack[-1]==x:
        stack.pop()
        result.append("-")
    else:
        ans = False
        break

if ans==False:
    print("NO")
else:
    for v in result:
        print(v)