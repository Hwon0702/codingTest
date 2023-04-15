import sys
input = sys.stdin.readline 
sys.setrecursionlimit(10**6)
n = int(input())
p = list(map(int, input().split()))
s = list(map(int, input().split()))
g = [0, 1, 2] * (n // 3)
origin = p
new = [0 for _ in range(n)]
cnt = 0
while p !=g:
    for i in range(n):
      new[s[i]]=p[i]
    p = new
    new = [0 for _ in range(n)]
    cnt+=1
    if origin ==p:
       cnt=-1
       break
print(cnt)
