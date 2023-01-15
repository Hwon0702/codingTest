import sys 
from collections import deque
input = sys.stdin.readline 
sys.setrecursionlimit(10**6)
n, m = map(int, input().split())
graph = [set()for _ in range(n+1)]
res = [999999 for _ in range(n+1)]
def find(s):
    q = deque()
    cost = [999 for _ in range(n+1)]
    visited = [False for _ in range(n+1)]
    visited[0]=True
    visited[s]=True
    q.append(s)
    cost[s]=0
    while q:
        now = q.popleft()
        for i in graph[now]:
            if not visited[i]:
                q.append(i)
                visited[i]=True
                cost[i]=min(cost[i], cost[now]+1)
    return sum(cost[1:])

for _ in range(m):
    s, e = map(int, input().split())
    graph[s].add(e)
    graph[e].add(s)
for i in range(1, n+1):
    res[i] =find(i)
print(res.index(min(res)))