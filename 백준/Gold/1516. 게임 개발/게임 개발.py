
from collections import deque
import sys
input = sys.stdin.readline
sys.setrecursionlimit(10**6)

n = int(input())
times = [[] for _ in range(n+1)]
align = [[] for _ in range(n+1)]
cnt = [0]+[0 for _ in range(n)]
ans = [0]+[0 for _ in range(n)]

res=[]
for i in range(1, n+1):
    m = list(map(int, input().split()))
    times[i]=m[0]
    for be in m[1:-1]:
        align[be].append(i)
        cnt[i]+=1

q = deque()
for i in range(1, n+1):
    if cnt[i] == 0: 
        q.append(i)
        ans[i]=times[i]
while q:
    node = q.popleft()
    for i in align[node]:
        cnt[i] -=1
        ans[i] = max(ans[i], times[i]+ans[node])
        if cnt[i] == 0:
            q.append(i)
for v in range(1, n+1):
    print(ans[v])
