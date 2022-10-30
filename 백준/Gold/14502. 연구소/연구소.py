import sys 
import copy
from collections import deque
input = sys.stdin.readline

def bfs():
    q = deque()
    tmp = copy.deepcopy(graph)
    for x in range(n):
        for y in range(m):
            if graph[x][y]==2:
                q.append([x,y])
    while q:
        cx, cy = q.popleft()
        for i in range(4):
            nx = cx+dx[i]
            ny = cy+dy[i]
            if 0<=nx<n and 0<=ny<m and tmp[nx][ny]==0:
                tmp[nx][ny]=2
                q.append([nx, ny])
    res = 0
    for i in range(n):
        res+=tmp[i].count(0)
    return res


def wall(cnt):
    global ans
    if cnt==3:
        a = bfs()
        ans = max(ans,a)
        return
    for x in range(n):
        for y in range(m):
            if graph[x][y]==0:
                graph[x][y]=1
                wall(cnt+1)
                graph[x][y]=0

n,m = map(int,input().split())
graph = []
for _ in range(n):
    graph.append(list(map(int, input().split())))
dx = [-1,1,0,0]
dy = [0,0,-1,1]
ans = 0
wall(0)
print(ans)