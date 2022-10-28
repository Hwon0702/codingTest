import sys 
input = sys.stdin.readline
sys.setrecursionlimit(10**6)

v, e = map(int, input().split())
graph = [[]for _ in range(v+1)]
reverse_graph = [[]for _ in range(v+1)]
for i in range(e):
    s, f = map(int, input().split())
    graph[s].append(f)
    reverse_graph[f].append(s)

def dfs(start, visited, stack):
    visited[start]=True
    for now in graph[start]:
        if visited[now]==False:
            dfs(now, visited, stack)
    stack.append(start)

def reverse_dfs(start, visited, stack):
    visited[start]=True
    stack.append(start)
    for now in reverse_graph[start]:
        if visited[now]==False:
            reverse_dfs(now, visited, stack)
            
stack = []
visited = [False for _ in range(v+1)]
for i in range(1, v+1):
    if visited[i]==False:
        dfs(i, visited, stack)
visited = [False for _ in range(v+1)]
res = []
while stack:
    scc = []
    now=stack.pop()
    if visited[now]==False:
        reverse_dfs(now, visited, scc)
        res.append(sorted(scc))
res = sorted(res)
print(len(res))
for v in res:
    print(*v, -1)