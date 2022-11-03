import sys 
from collections import deque
input = sys.stdin.readline
sys.setrecursionlimit(10**6)

def bfs(start, connected):
    q = deque()
    q.append(start)
    visited[start]=connected
    while q:
        now = q.popleft()
        for i in graph[now]:
            if visited[i]==False:
                q.append(i)
                visited[i]=not visited[now] #다른 그룹으로 지정
            elif visited[i]==visited[now]:
                return False
    return True

K = int(input()) #TC
for _ in range(K):
    res = True
    V, E = map(int, input().split())
    graph = [[] for _ in range(V)]
    visited = [False for _ in range(V)]
    for i in range(E):
        u, v = map(int, input().split())
        graph[u-1].append(v-1)
        graph[v-1].append(u-1)
    
    for i in range(V):
        if visited[i]==False:
            res = bfs(i, 0)
            if res ==False:
                break
    if res:
        print("YES")
    else:
        print("NO")