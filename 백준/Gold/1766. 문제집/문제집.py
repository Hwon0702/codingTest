import sys 
import heapq
from collections import deque
input = sys.stdin.readline
sys.setrecursionlimit(10**6)
n, m = map(int, input().split())
problems = [[]for _ in range(n+1)]
res = []
cnt = [0 for _ in range(n+1)]
q = []
for _ in range(m):
    start, end = map(int, input().split())
    problems[start].append(end)
    cnt[end]+=1

for i in range(1, n+1):
    if cnt[i]==0:
        heapq.heappush(q, i)
while q:
    now =heapq.heappop(q)
    res.append(now)
    for v in problems[now]:
        cnt[v]-=1
        if cnt[v]==0:
            heapq.heappush(q, v)
print(' '.join(map(str, res)))