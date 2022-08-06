from collections import deque
import re
import sys
input = sys.stdin.readline
sys.setrecursionlimit(10**6)
from collections import deque
def bfs():
    while q:
        x, y = q.popleft()
        for i in range(4):
            nx = x+dx[i]
            ny = y+dy[i]
            if 0<=nx<h and 0<=ny<w and graph[nx][ny]==0 :
                graph[nx][ny]=graph[x][y]+1
                q.append([nx, ny])
                

w, h = map(int, input().split())
graph = []
dx = [-1, 1, 0, 0]
dy = [0, 0,  -1, 1]
q = deque()
for i in range(h):
    line = list(map(int, input().split()))
    graph.append(line)
    for j in range(w):
        if graph[i][j]==1:
            q.append([i,j])
bfs()
res = 0
impossible = False
for i in range(h):
    if impossible:
        break
    for j in range(w):
        if graph[i][j]==0:
            impossible=True
            break
        else:
            res=max(res, graph[i][j])
if impossible:
    print(-1)
else:
    print(res-1)
