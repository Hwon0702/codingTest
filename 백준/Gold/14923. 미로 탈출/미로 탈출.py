import sys 
from collections import deque
input = sys.stdin.readline
sys.setrecursionlimit(10**6)
n, m = map(int, input().split())
hx, hy = map(int, input().split())
ex, ey = map(int, input().split())
graph = []
dx = [-1, 1, 0, 0]
dy = [0, 0, -1, 1]
for _ in range(n):
    graph.append(list(map(int, input().split())))
hx, hy, ex, ey = hx-1, hy-1, ex-1, ey-1
def find():
    queue = deque()
    queue.append([hx,hy,0,1])
    visited = [[[False for _ in range(m)] for _ in range(n)] for _ in range(2)]
    visited[1][hx][hy] = True
    while queue:
        x,y,cnt,used = queue.popleft()
        if x==ex and y==ey:
            return cnt
        for i in range(4):
            nx = x+dx[i]
            ny = y+dy[i]
            if 0<=nx<n and 0<=ny<m:
                if used:
                    if not graph[nx][ny]:
                        if not visited[1][nx][ny]:
                            visited[1][nx][ny] = True
                            queue.append([nx,ny,cnt+1,used])
                    elif graph[nx][ny]:
                        if not visited[0][nx][ny]:
                            visited[0][nx][ny] = True
                            used = 0
                            queue.append([nx,ny,cnt+1,used])
                            used = 1
                elif not used:
                    if not visited[0][nx][ny]:
                        if not graph[nx][ny]:
                            visited[0][nx][ny] = True
                            queue.append([nx,ny,cnt+1,used])
    return -1

print(find())