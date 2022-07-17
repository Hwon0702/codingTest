import sys
import collections
input = sys.stdin.readline
sys.setrecursionlimit(10**6)
n, m = map(int, input().split())
maze = [list(map(int, ' '.join(input()).split())) for _ in range(n)]

dx = [1, -1, 0, 0]
dy = [0, 0, 1, -1]

queue = collections.deque([(0, 0)]) #시작노드
result = 0

while queue:
    x, y = queue.popleft()
    for i in range(4):
        nx, ny = x + dx[i], y + dy[i] #상하좌우이동
        if 0 <= nx < n and 0 <= ny < m:
            if maze[nx][ny] == 1:
                maze[nx][ny] = maze[x][y] + 1
                queue.append((nx, ny))

print(maze[n - 1][m - 1])
