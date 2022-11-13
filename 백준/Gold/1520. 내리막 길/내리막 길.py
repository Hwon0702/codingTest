import sys 
input = sys.stdin.readline
sys.setrecursionlimit(10**6)
m, n =map(int, input().split())
graph = []
routes = [[-1 for _ in range(n)]for _ in range(m)]
for _ in range(m):
    graph.append(list(map(int, input().split())))
dx = [-1,1,0,0]
dy = [0,0,-1,1]
def dfs(x,y):
    if x==m-1 and y==n-1:
        return 1
    if routes[x][y]>=0:
        return routes[x][y]
    routes[x][y]=0
    for i in range(4):
        nx = x+dx[i]
        ny = y+dy[i]
        if 0<=nx<m and 0<=ny<n and graph[nx][ny]<graph[x][y]:
            routes[x][y]+=dfs(nx, ny)
    
    return routes[x][y]

print(dfs(0,0))