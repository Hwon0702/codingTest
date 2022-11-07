from collections import deque
from itertools import combinations
import sys 
input = sys.stdin.readline
sys.setrecursionlimit(10**6)
n, m = map(int, input().split())
dx = [-1,1,0,0]
dy = [0,0,-1,1]
graph = []
virus = []
res = float('inf')
for i in range(n):
    tmp = list(map(int, input().split()))
    for j in range(n):
        if tmp[j]==2:
            virus.append([i, j])
    graph.append(tmp)
def bfs(virus_list):
    actived = 0
    q = deque()
    visited = [[-1 for _ in range(n)]for _ in range(n)]
    for x, y in virus_list:
        q.append([x, y])
        visited[x][y]=0
    while q:
        cx, cy = q.popleft()
        for i in range(4):
            nx = cx+dx[i]
            ny = cy+dy[i]
            if 0 <= nx < n and 0 <= ny < n and visited[nx][ny] == -1:
                if graph[nx][ny]==0:
                    visited[nx][ny]=visited[cx][cy]+1
                    actived = max(actived, visited[nx][ny])
                    q.append([nx, ny])
                elif graph[nx][ny]==2:
                    visited[nx][ny]=visited[cx][cy]+1
                    q.append([nx, ny])

    for i in range(n):
        for j in range(n):
            if graph[i][j]==0 and visited[i][j]<0:
                return float('inf')
    return actived
    

for active in combinations(virus, m):
    res = min(res, bfs(active))
print(res if res != float('inf') else -1)