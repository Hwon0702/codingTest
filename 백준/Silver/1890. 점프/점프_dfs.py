import sys 
from collections import deque
input = sys.stdin.readline
sys.setrecursionlimit(10**6)
n = int(input())
graph = [list(map(int, input().split())) for _ in range(n)]
q = deque()
q.append((0, 0))
cnt = 0
while q:
    now_x, now_y = q.popleft()
    if now_x ==n-1 and now_y ==n-1:
        cnt+=1
        continue
    elif 0<=now_x<n and 0<=now_y<n:
        if 0<=now_x+graph[now_x][now_y]<n:
            nx = now_x+graph[now_x][now_y]
            q.append((nx, now_y))
        if 0<=now_y+graph[now_x][now_y]<n:
            ny = now_y+graph[now_x][now_y]
            q.append((now_x, ny))
print(cnt)