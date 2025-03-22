import sys 
from collections import deque
input = sys.stdin.readline 
sys.setrecursionlimit(10**6)

def bfs(x, y):
    graph[x][y] = 0
    dx = [1, -1, 0, 0]
    dy = [0, 0, 1, -1]
    width = 1
    q = deque()
    q.append([x, y])
    while q:
        x, y = q.popleft()
        for i in range(4):
            nx = dx[i] + x
            ny = dy[i] + y
            if 0 <= nx < n and 0 <= ny < m and graph[nx][ny] == 1:
                q.append([nx, ny])
                graph[nx][ny] = 0
                width += 1
    return width


n, m = map(int, input().split())
graph = [list(map(int, input().split())) for _ in range(n)]
cnt = 0
res = 0
for i in range(n):
    for j in range(m):
        if graph[i][j] == 1:
            cnt += 1
            res = max(bfs(i, j), res)
print(cnt)
print(res)