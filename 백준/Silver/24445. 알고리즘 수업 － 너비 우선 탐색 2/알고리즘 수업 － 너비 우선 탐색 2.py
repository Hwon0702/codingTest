import sys
from collections import deque
input = sys.stdin.readline
sys.setrecursionlimit(10**6)
n,m,r = map(int,input().split())
graph = [[] for _ in range(n+1)]
visited = [0] * (n+1)
visited[r] = 1
cnt = 1
q = deque()
q.append(r)

for _ in range(m):
    a,b = map(int,input().split())
    graph[a].append(b)
    graph[b].append(a)

for i in range(1,n+1):
    graph[i].sort(reverse = True)

while q:
    v = q.popleft()
    for i in graph[v]:
        if visited[i]==False:
            cnt+=1
            visited[i] = cnt
            q.append(i)
for v in visited[1:]:
    print(v)
