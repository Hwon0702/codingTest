import sys 
input = sys.stdin.readline 
sys.setrecursionlimit(10**6)
from collections import deque
dx = [-1,1,0,0]
dy = [0,0,-1,1]
arr = []
res = 0
def bfs(start_x, start_y):
    q = deque()
    visited = [[False for _ in range(c)]for _ in range(r)]
    q.append([start_x, start_y])
    sheep = 0 
    wolf = 0
 
    if graph[start_x][start_y] == 'o':
        sheep += 1
    elif graph[start_x][start_y] == 'v':
        wolf += 1
    graph[start_x][start_y] = '#'
    while q:
        cx, cy = q.popleft()
        for i in range(4):
            nx = cx+dx[i]
            ny = cy+dy[i]
            if 0<=nx<r and 0<=ny<c and visited[nx][ny]==False and graph[nx][ny]!='#':
                if graph[nx][ny] == 'o':
                    sheep += 1
                elif graph[nx][ny] == 'v':
                    wolf += 1
                graph[nx][ny] = '#'
                q.append([nx, ny])
    
    if wolf >= sheep:
        return 0, wolf
    elif sheep > wolf:
        return sheep, 0
    
r, c = map(int, input().split())
graph = []
total_sheep = 0
total_wolf = 0
for i in range(r):
    graph.append(list(map(str, input().rstrip('\n'))))


for i in range(r):
    for j in range(c):
        if graph[i][j] =='o'or graph[i][j] =='v':
            sheep, wolf = bfs(i, j)
            total_sheep += sheep
            total_wolf += wolf
print(total_sheep, total_wolf)