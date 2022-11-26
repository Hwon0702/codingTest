from collections import deque
import sys
input = sys.stdin.readline
n = int(input())

graph = [list(map(int, input().split())) for _ in range(n)]
visited = [[False for _ in range(n)]for _ in range(n)]
dx = [-1,1,0,0]
dy = [0,0,-1,1]

def devide_island(i, j):
    island = deque()
    island.append([i, j])
    while island:
        cx, cy = island.popleft()
        for i in range(4):
            nx = cx+dx[i]
            ny = cy+dy[i]
            if 0<=nx<n and 0<=ny<n and visited[nx][ny]==False and graph[nx][ny]==1:
                visited[nx][ny] = True
                graph[nx][ny] = island_num
                island.append([nx,ny])

def making_bridge(start_island):
    q = deque()
    length = [[float('inf') for _ in range(n)] for _ in range(n)]
    for i in range(n):
        for j in range(n):
            if graph[i][j]==start_island:
                length[i][j] = 0
                q.append([i,j])
    while q:
        cx, cy = q.popleft()
        for i in range(4):
            nx = cx+dx[i]
            ny = cy+dy[i]
            if 0<=nx<n and 0<=ny<n:
                if graph[nx][ny]!=0 and graph[nx][ny]!=start_island and visited[nx][ny]:
                    return length[cx][cy]
                elif graph[nx][ny]==0 and length[nx][ny]==float('inf'):
                    length[nx][ny] = length[cx][cy]+1
                    q.append([nx,ny])
    return float('inf')
island_num=1
for i in range(n):
    for j in range(n):
        if graph[i][j]==1 and visited[i][j]==False:
            visited[i][j] = True
            graph[i][j] = island_num
            devide_island(i,j)
            island_num += 1

res = float('inf')
for i in range(1,island_num):
    res = min(res, making_bridge(i))
print(res)