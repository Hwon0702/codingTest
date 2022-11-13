from collections import deque
import sys 
input = sys.stdin.readline
def bfs(start_x, start_y):
    visited = [[False for _ in range(w)]for _ in range(h)]
    visited[start_x][start_y] = True
    q = deque()
    q.append([start_x, start_y])
    max_cost = 0
    while q:
        cx, cy = q.popleft()
        for i in range(4):
            nx = cx+dx[i]
            ny = cy+dy[i]
            if 0<=nx<h and 0<=ny<w and visited[nx][ny]==False and graph[nx][ny]=='L':
                visited[nx][ny]=True
                q.append([nx, ny])
                res[nx][ny] = res[cx][cy]+1
                max_cost = max(max_cost, res[nx][ny])
    return max_cost

h, w = map(int, input().split())
graph = []

res = [[0 for _ in range(w)]for _ in range(h)]
result=0
dx = [-1, 1, 0, 0]
dy = [0, 0, -1, 1]
for i in range(h):
    graph.append(list(input().strip('\n')))

for i in range(h):
    for j in range(w):
        if graph[i][j]=='L':
            res[i][j] = 0
            result = max(result, bfs(i, j))
print(result)