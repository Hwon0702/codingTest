import sys
sys.setrecursionlimit(10**6)
input = sys.stdin.readline
def dfs(v):
    global cnt 
    visited[v]=True
    res[v] = cnt
    graph[v].sort(reverse=True)
    for i in graph[v]:
        if not visited[i]:
            cnt += 1
            dfs(i)

n, m, s = map(int, input().split())
graph = [[] for _ in range(n+1)]
visited = [False] * (n+1)
res = [0] * (n+1)
cnt=1
for _ in range(m):
    a, b = map(int, input().split())
    graph[a].append(b)
    graph[b].append(a)

dfs(s)
for v in res[1:]:
    print(v)