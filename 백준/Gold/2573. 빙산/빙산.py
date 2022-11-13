import sys 
from collections import deque
input = sys.stdin.readline
dx = [-1,1,0,0]
dy = [0,0,-1,1]
day = 0
split =False
n, m = map(int, input().split())
graph = []
res = []
for _ in range(n):
    graph.append(list(map(int, input().split())))
def bfs(x, y):
    q=deque()
    q.append([x,y])
    visited[x][y]=True
    while q:
        cx, cy = q.popleft()
        for i in range(4):
            nx = cx+dx[i]
            ny = cy+dy[i]
            if 0<=nx<n and 0<=ny<m:
                if visited[nx][ny]==False and graph[nx][ny]>0:
                    visited[nx][ny]=True
                    q.append([nx, ny])
                elif graph[nx][ny]==0:
                    count[cx][cy]+=1
    return 1

while True:
    visited = [[False for _ in range(m)]for _ in range(n)]
    count = [[0 for _ in range(m)]for _ in range(n)]
    res = []
    for i in range(n):
        for j in range(m):
            if graph[i][j]>0 and visited[i][j]==False:
                res.append(bfs(i, j))
    for i in range(n):
        for j in range(m):
            graph[i][j]-=count[i][j]
            if graph[i][j]<0:
                graph[i][j]=0
    if len(res) ==0:
        break
    if len(res)>=2:
        split = True
        break
    day+=1
print(day if split else 0)