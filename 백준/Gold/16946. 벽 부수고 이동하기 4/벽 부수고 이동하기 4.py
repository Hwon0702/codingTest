import sys 
from collections import deque
input = sys.stdin.readline 
sys.setrecursionlimit(10**6)

n, m= map(int, input().split())
dx = [-1,1,0,0]
dy = [0,0,-1,1]
graph=[]
for _ in range(n):
    graph.append(list(map(str, input().rstrip('\n'))))
visited=[[False for _ in range(m)] for _ in range(n)]
res=[[0 for _ in range(m)] for _ in range(n)]

for i in range(n):
    for j in range(m):
        if graph[i][j] == '1': 
            res[i][j] = 1
def bfs(i, j):
    q = deque()
    q.append([i, j])
    visited[i][j]=True
    cnt = 1
    next_q = []
    while q:
        cx, cy = q.popleft()
        for i in range(4):
            nx = cx+dx[i]
            ny = cy+dy[i]
            if 0<=nx<n and 0<=ny<m and visited[nx][ny]==False:
                visited[nx][ny]=True
                if graph[nx][ny]=='0':
                    visited[nx][ny]=True
                    q.append([nx, ny])
                    cnt+=1
                else:
                    next_q.append([nx, ny])
    for x, y in next_q:
        visited[x][y] = False
        res[x][y] += cnt
        if res[x][y] >= 10: 
            res[x][y] %= 10

for i in range(n):
    for j in range(m):
        if graph[i][j] == '0':
            if not visited[i][j]:
                visited[i][j] = True
                bfs(i,j)

for i in range(n):
    print(''.join(map(str,res[i])))