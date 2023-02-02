import sys 
from collections import deque
input = sys.stdin.readline
sys.setrecursionlimit(10**6)

n, m = map(int, input().split())
cheese = []
time = 0
dx = [-1,1,0,0]
dy = [0,0,-1,1]

for _ in range(n):
    cheese.append(list(map(int, input().split())))

def melt():
    q = deque()
    q.append([0,0])
    global visited
    visited[0][0]=True 
    while q:
        cx, cy = q.popleft()
        cnt= 0
        for i in range(4):
            nx = cx+dx[i]
            ny = cy+dy[i]
            if 0<=nx<n and 0<=ny<m and visited[nx][ny]==False:
                if cheese[nx][ny]>=1:
                    cheese[nx][ny]+=1
                else:
                    visited[nx][ny]=True
                    q.append([nx, ny])

    return 
    
while 1:
    break_flag = 0
    visited = [[False for _ in range(m)]for _ in range(n)]
    melt()
    for i in range(n):
        for j in range(m):
            if cheese[i][j]>=3:
                cheese[i][j]=0
                break_flag = 1
            elif cheese[i][j]==2:
                cheese[i][j]=1
    if break_flag>0:
        time+=1
    else:
        break
print(time)
