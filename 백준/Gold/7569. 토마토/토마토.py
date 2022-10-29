import sys 
from collections import deque
input = sys.stdin.readline
sys.setrecursionlimit(10**6)
m, n, h= map(int, input().split())
q= deque()
graph = []
for i in range(h):
    tmp = []
    for j in range(n):
        tmp.append(list(map(int,input().split())))
        for k in range(m):
            if tmp[j][k]==1:
                q.append([i,j,k])
    graph.append(tmp)

dx = [-1,1,0,0,0,0]
dy = [0,0,1,-1,0,0]
dz = [0,0,0,0,1,-1]

while(q):
    cx, cy, cz = q.popleft()
    
    for i in range(6):
        nx = cx+dx[i]
        ny = cy+dy[i]
        nz = cz+dz[i]
        if 0<=nx<h and 0<=ny<n and 0<=nz<m and graph[nx][ny][nz]==0:
            q.append([nx, ny, nz])
            graph[nx][ny][nz] = graph[cx][cy][cz]+1
day = 0
for i in graph:
    for j in i:
        for k in j:
            if k==0:
                print(-1)
                exit(0)
            day = max(day,k)
print(day-1)
 