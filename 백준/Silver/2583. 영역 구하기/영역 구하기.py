import sys
from collections import deque
input = sys.stdin.readline 
sys.setrecursionlimit(10**6)
m, n, k = map(int, input().split())
excepts = []
graph = [[0 for _ in range(n)] for _ in range(m)]

dx = [-1, 1, 0, 0]
dy = [0, 0, -1, 1]
res = []
for _ in range(k):
    sx, sy, ex, ey = map(int, input().split())
    for i in range(sx, ex):
        for j in range(sy, ey):
            graph[j][i]=1 #건드릴 수 없는 영역 처리
def fill(x, y):
    q = deque()
    q.append([x, y])
    cnt = 1
    global graph
    while q:
        cx, cy = q.popleft()
        for i in range(4):
            nx = cx+dx[i]
            ny = cy+dy[i]
            if 0<=nx<m and 0<=ny<n and graph[nx][ny]==0:
                q.append([nx, ny])
                graph[nx][ny]=1
                cnt+=1 #영역의 넓이
    return cnt
for i in range(m):
    for j in range(n):
        if graph[i][j] == 0:
            graph[i][j] = 1
            res.append(fill(i, j))
print(len(res))
res.sort()
print(*res)