import sys 
from collections import deque
input = sys.stdin.readline
sys.setrecursionlimit(10**6)
dx = [0, 0, -1, 1]
dy = [1, -1, 0, 0]

h, w= map(int, input().split())
graph = [list(map(str, input())) for _ in range(h)]
def bfs():
    visited = [[[0]*2 for _ in range(w)] for _ in range(h)]
    queue = deque()
    queue.append((0,0,0))
    visited[0][0][0] = 1

    while queue:
        x, y, z = queue.popleft()
        if x==h-1 and y==w-1:
            return visited[x][y][z]
        for i in range(4):
            nx = x+dx[i]
            ny = y+dy[i]
            if 0<=nx<h and 0<=ny<w and visited[nx][ny][z]==0:
                if graph[nx][ny]=='0':
                    visited[nx][ny][z]=visited[x][y][z]+1
                    queue.append((nx,ny,z))
                elif graph[nx][ny]=='1' and z==0:
                    visited[nx][ny][z+1]=visited[x][y][z]+1
                    queue.append((nx,ny,z+1))

    return -1
print(bfs())