from collections import deque
import sys 
input = sys.stdin.readline 
sys.setrecursionlimit(10**6)
dx = [0, 1, 0, -1]
dy = [1, 0, -1, 0]
n = int(input())
graph = [[0 for _ in range(n)] for _ in range(n)]
graph[0][0]=1#뱀시작
apple = int(input())
dist = deque()
for _ in range(apple):
    x, y = map(int, input().split())
    x, y = x-1, y-1
    graph[x][y]=2 #사과
d = int(input())
for _ in range(d):
    a, b = map(str, input().split())
    dist.append([int(a), b]) #방향전환
cnt = 0
q = deque()
snake = deque()
snake.append([0,0])#뱀 경로
q.append([0, 0, 0])#끝, dist index
while True:
    cnt+=1
    cx, cy, d = q.popleft()
    nx = cx+dx[d]
    ny = cy+dy[d]
    if not(0<=nx<n and 0<=ny<n) or [nx, ny] in snake:
        break
    if 0<=nx<n and 0<=ny<n:
        snake.append([nx,ny])#뱀 경로
        if graph[nx][ny] != 2:
            sx, sy = snake.popleft()
            graph[sx][sy] =0
        
        graph[nx][ny]=1
    if dist and cnt == dist[0][0]:
        if dist[0][1] == 'L':
            d = (d - 1) % 4
        else:
            d = (d + 1) % 4
        dist.popleft()
    q.append([nx, ny, d])

print(cnt)