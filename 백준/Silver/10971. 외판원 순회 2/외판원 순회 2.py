import sys 
input = sys.stdin.readline
sys.setrecursionlimit(10**6)
n = int(input())
graphs = []
visited = [[False for _ in range(n)] for _ in range(n)]

def dfs(start, now, cost, cnt):
    global res
    if cnt == n:
        if graphs[now][start]:
            cost += graphs[now][start]
            if res > cost:
                res = cost
        return

    if cost > res:
        return

    for i in range(n):
        if not visited[i] and graphs[now][i]:
            visited[i] = 1
            dfs(start, i, cost + graphs[now][i], cnt + 1)
            visited[i] = 0

    
for _ in range(n):
    graphs.append(list(map(int, input().split())))
res = sys.maxsize
for i in range(n):
    visited[i]=True
    dfs(i, i, 0, 1)
    visited[i]=False
print(res)