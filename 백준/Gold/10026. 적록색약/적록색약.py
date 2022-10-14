import sys 
from collections import deque
input = sys.stdin.readline
sys.setrecursionlimit(10**6)
n= int(input())
pic = [list(str(input().strip('\n'))) for _ in range(n)]
dx = [0,0,-1,1]
dy = [-1,1,0,0]

def bfs(x, y):
    q.append([x, y])
    while q:
        now_x, now_y = q.popleft()
        for i in range(4):
            nx = now_x+dx[i]
            ny = now_y+dy[i]
            if 0<=nx<n and 0<=ny<n and not visited[nx][ny]:
                if pic[nx][ny] == pic[now_x][now_y]:
                    visited[nx][ny] = True
                    q.append([nx, ny])
visited = [[False for _ in range(n)]for _ in range(n)]
q = deque()
normal = 0
for i in range(n):
    for j in range(n):
        if not visited[i][j]:
            bfs(i, j)
            normal+=1
visited = [[False for _ in range(n)]for _ in range(n)]
q = deque()
abnormal=0
for i in range(n):
    for j in range(n):
        if pic[i][j]=='G':
            pic[i][j]='R'

for i in range(n):
    for j in range(n):
        if not visited[i][j]:
            bfs(i, j)
            abnormal+=1

print(normal, abnormal)