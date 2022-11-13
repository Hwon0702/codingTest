import sys 
from collections import deque
input = sys.stdin.readline
sys.setrecursionlimit(10**6)
def dfs(x, y, w):
    q = deque()
    q.append((x, y))
    while q:
        x, y = q.popleft()
        for i in range(4):
            nx = x+dx[i]
            ny = y+dy[i]
            if (0<=nx<n)and(0<=ny<n):
                if graph[x][y]>w and visited[nx][ny]==0:
                    visited[nx][ny]=1
                    dfs(nx, ny, w)

n = int(input())
graph = [[] for _ in range(n)]
water = 0
safety_area = []
dx=[1,-1,0,0]
dy=[0,0,1,-1]
for i in range(n):
    graph[i]=list(map(int, input().split()))
    water = max(water, max(graph[i]))
for w in range(water):
    visited = [[0] * n for _ in range(n)]
    cnt = 0
    for i in range(n): 
        for j in range(n):
            if graph[i][j] > w and visited[i][j] == 0:
                dfs(i, j, w)
                cnt += 1
    safety_area.append(cnt)
print(max(safety_area))