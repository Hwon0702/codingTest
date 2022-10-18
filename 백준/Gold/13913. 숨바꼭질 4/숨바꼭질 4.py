import sys 
from collections import deque
input = sys.stdin.readline
sys.setrecursionlimit(10**6)
n, k = map(int, input().split())
visited = [False for _ in range(100001)]
route = [0 for _ in range(100001)]
q=deque()
q.append([n, 0])
route[n] = n
res = 0
while q:
    now, cnt = q.popleft()
    if now==k:
        res = cnt
        break
    if 0<=now*2<100001 and not visited[now*2]:
        visited[now*2]=True 
        route[now*2] = now
        q.append([now*2, cnt+1])
    if 0<=now-1<100001 and not visited[now-1]:
        visited[now-1]=True 
        route[now-1] = now
        q.append([now-1, cnt+1])
    if 0<=now+1<100001 and not visited[now+1]:
        visited[now+1]=True 
        route[now+1] = now
        q.append([now+1, cnt+1])

route_res = []
idx = k
for _ in range(res):
    route_res.append(idx)
    idx = route[idx]
route_res.append(n)
print(res)
print(' '.join(map(str, route_res[::-1])))