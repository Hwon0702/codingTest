import sys
from collections import deque
input=sys.stdin.readline
sys.setrecursionlimit(10**6)
def bfs(s, e):
    queue = deque()
    queue.append((s,0))
    visited = [False] * (n+1)
    visited[s] = True
    while queue:
        s, d = queue.popleft()
        if s == e:
            return d
        for i, l in graph[s]:
            if not visited[i]:
                visited[i] = True
                queue.append((i,d+l))  #시작할곳, 여태까지의 거리

n, m = map(int, input().split())
graph = [[] for _ in range(n+1)]
for _ in range(n-1):
    s, e, d = map(int,input().split())
    graph[s].append((e,d)) #양방향
    graph[e].append((s,d))
for i in range(m):
    s, e = map(int, input().split())
    print(bfs(s,e))