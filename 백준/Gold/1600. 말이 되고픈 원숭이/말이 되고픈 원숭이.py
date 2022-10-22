import sys 
from collections import deque
input = sys.stdin.readline
sys.setrecursionlimit(10**6)
dx = [0, 0, -1, 1]
dy = [1, -1, 0, 0]
hx = [-2, -2, -1, 1, 2, 2, -1, 1]
hy = [-1, 1, -2, -2, -1, 1, 2, 2]
k = int(input())
w, h= map(int, input().split())
graph = [list(map(int, input().split())) for _ in range(h)]
def bfs():
    visited = [[[0]*(k+1) for _ in range(w)] for _ in range(h)]
    queue = deque()
    queue.append((0,0,0))
    visited[0][0][0] = 1

    while queue:
        x, y, z = queue.popleft()
        if x==h-1 and y==w-1:
            return visited[x][y][z]-1
        for i in range(4):
            nx = x+dx[i]
            ny = y+dy[i]
            if 0<=nx<h and 0<=ny<w and not visited[nx][ny][z] and not graph[nx][ny]:
                visited[nx][ny][z] = visited[x][y][z]+1
                queue.append((nx,ny,z))
        if z<k:
            for i in range(8):
                nx= x+hx[i]
                ny = y+hy[i]
                if 0<=nx<h and 0<=ny<w:
                    if not graph[nx][ny]:
                        if not visited[nx][ny][z+1]:
                            visited[nx][ny][z+1] = visited[x][y][z]+1
                            queue.append((nx,ny,z+1))

    return -1
print(bfs())