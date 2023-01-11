import sys 
from collections import deque
input = sys.stdin.readline
sys.setrecursionlimit(10**6)
dx = [-2, -2, -1, -1, 1, 1, 2, 2]
dy = [-1, 1, -2, 2, -2, 2, -1, 1]

def search(now, target, width):
    if now==target:
        return 0
    q = deque()
    visited = [[False for _ in range(width)] for _ in range(width)]
    q.append([0, now[0], now[1]])
    visited[now[0]][now[1]]=True
    while q:
        cnt, cx, cy = q.popleft()
        if cx ==target[0] and cy==target[1]:
            return cnt
        for i in range(8):
            nx = cx+dx[i]
            ny = cy+dy[i]
            if 0<=nx<width and 0<=ny<width and not visited[nx][ny]:
                q.append([cnt+1, nx, ny])
                visited[nx][ny]=True

tc = int(input())
for _ in range(tc):
    L = int(input())
    now = list(map(int, input().split()))
    target = list(map(int, input().split()))
    print(search(now, target, L))