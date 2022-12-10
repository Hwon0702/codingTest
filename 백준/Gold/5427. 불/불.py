from collections import deque
import sys
input = sys.stdin.readline
sys.setrecursionlimit(10**6)
dx = [-1,1,0,0]
dy = [0,0,-1,1]


def fire_spread():
    for _ in range(len(fire)):
        x, y = fire.popleft()
        for i in range(4):
            nx = x + dx[i]
            ny = y + dy[i]
            if 0 <= nx < h and 0 <= ny < w:
                if arr[nx][ny] != '#' and arr[nx][ny] != '*':
                    arr[nx][ny] = '*'
                    fire.append((nx, ny))
def person_move():
    possible = False
    for _ in range(len(start)):
        x, y = start.popleft()
        for i in range(4):
            nx = x + dx[i]
            ny = y + dy[i]
            if 0 <= nx < h and 0 <= ny < w:
                if visited[nx][ny] == 0 and arr[nx][ny] != '#' and arr[nx][ny] != '*':
                    possible = True
                    visited[nx][ny] = visited[x][y] + 1
                    start.append((nx, ny))
            else:
                return visited[x][y]
    if not possible:
        return 'IMPOSSIBLE'


T = int(input())
for _ in range(T):
    w, h = map(int, input().split())
    arr = []
    fire = deque()
    start = deque()
    for i in range(h):
        arr.append(list(input().strip('\n')))
        for j in range(w):
            if arr[i][j] == '*':
                fire.append((i, j))
            if arr[i][j] == '@':
                start.append((i, j))
    visited = [[0] * w for _ in range(h)]
    visited[start[0][0]][start[0][1]] = 1

    result = 0
    while True:
        fire_spread()
        result = person_move()
        if result:
            break

    print(result)