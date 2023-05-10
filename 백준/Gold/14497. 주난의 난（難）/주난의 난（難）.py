import sys 
from collections import deque

input = sys.stdin.readline 
sys.setrecursionlimit(10**6)
dx = [-1,1,0,0]
dy = [0,0,-1,1]
def find(start_x, start_y):
    q = deque()
    global room
    visited = [[-1 for _ in range(m)]for _ in range(n)]
    q.append([start_x, start_y])
    visited[start_x][start_y]=0
    while q:
        cx, cy = q.popleft()
        if room[cx][cy]=='#':
            return visited[cx][cy]
        for i in range(4):
            nx = cx+dx[i]
            ny = cy+dy[i]
            if 0 <= nx < n and 0 <= ny < m and visited[nx][ny] == -1:
                if room[nx][ny] == "1" or room[nx][ny] == "#":
                    visited[nx][ny] = visited[cx][cy] + 1
                    q.append([nx, ny])
                elif room[nx][ny] == "0":
                    visited[nx][ny] = visited[cx][cy]
                    q.appendleft([nx, ny])
    return -1

n, m = map(int, input().split())
x1, y1, x2, y2 = map(int, input().split())
room = []
for _ in range(n):
    room.append(list(input().rstrip('\n')))
print(find(x1-1, y1-1))
#0은 빈 공간, 1은 친구, *은 주난, #은 범인
