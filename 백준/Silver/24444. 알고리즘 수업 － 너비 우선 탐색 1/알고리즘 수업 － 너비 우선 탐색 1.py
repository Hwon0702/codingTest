import sys 
from collections import deque
input = sys.stdin.readline 
sys.setrecursionlimit(10**6)

def bfs(start):
    visited=[False for _ in range(n+1)]
    visited[start]=True
    rating = [0 for _ in range(n+1)]
    cnt = 1
    q = deque()
    q.append(start)
    
    while q:
        now = q.popleft()
        rating[now]=cnt
        cnt+=1
        for i in graph[now]:
            if visited[i]==False:
                q.append(i)
                visited[i]=True
    for i in range(1, n+1):
        print(rating[i])

n, m, r = map(int, input().split())
graph = [[]for _ in range(n+1)]
for _ in range(m):
    s, e = map(int, input().split())
    graph[s].append(e)
    graph[e].append(s)
for i in range(1, n+1):
    graph[i].sort()
bfs(r)