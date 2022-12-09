from collections import deque
import sys
input = sys.stdin.readline
sys.setrecursionlimit(10**6)
n, k, m = map(int, input().split())

graph = [[] for _ in range(n+m+1)]
for i in range(1, m+1):
    sub = list(map(int, input().split()))
    graph[n+i] = sub
    for s in sub:
        graph[s].append(n+i)

q = deque()
visited = [False for _ in range(n+m+1)]
q.append((1, visited, 1))
visited[1] = True
res = 0
while q:
    subway, visited, cnt = q.popleft()
    if subway == n:
        res = cnt
        break

    for s in graph[subway]:
        if visited[s] == False:
            visited[s] = True
            if s > n:
                q.append((s, visited, cnt))
            else:
                q.append((s, visited, cnt+1))
if res>0:
    print(res)
else:
    print(-1)