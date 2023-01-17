import sys 
from collections import deque
input = sys.stdin.readline
sys.setrecursionlimit(10**6)

def air():
    visited = [[False for _ in range(m)]for _ in range(n)]
    cnt = 0
    q = deque()
    visited[0][0]=True
    q.append([0,0])
    while q:
        cx, cy = q.popleft()
        for i in range(4):
            nx = cx+dx[i]
            ny = cy+dy[i]
            if 0<=nx<n and 0<=ny<m and not visited[nx][ny]:
           
                if cheese[nx][ny]==0:
                    visited[nx][ny]=True
                    q.append([nx, ny])
                elif cheese[nx][ny]==1:
                    cheese[nx][ny]=0
                    cnt+=1
                    visited[nx][ny]=True
    res.append(cnt)
    return cnt

n, m = map(int, input().split())
cheese = []
res = []
time = 0
dx = [-1,1,0,0]
dy = [0,0,-1,1]

for i in range(n):
    cheese.append(list(map(int, input().split())))

while True:
    time+=1
    cnt = air()
    if cnt ==0:
        break
print(time-1)
print(res[-2])