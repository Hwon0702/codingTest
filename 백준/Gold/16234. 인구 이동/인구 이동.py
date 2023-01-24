import sys 
from collections import deque
input = sys.stdin.readline 
sys.setrecursionlimit(10**6)
n, l, r = map(int, input().split())
#인구의 차이가 l<=r
world = []
day = 0 #인구이동 일어나는 날
dx = [-1, 1, 0, 0]
dy = [0, 0, -1, 1]
#국경선을 공유해야함, 모든 왕국에 대해서 확인 필요
for _ in range(n):
    world.append(list(map(int, input().split())))

def bfs(a, b):
    united = deque()
    united_people, united_num = 0, 0
    visited[a][b] = 1
    while q:
        cx, cy = q.popleft()
        united.append([cx, cy])
        united_num += 1
        united_people += world[cx][cy]
        for i in range(4):
            nx = cx + dx[i]
            ny = cy + dy[i]
            if 0 <= nx < n and 0 <= ny < n:
                if visited[nx][ny] == 0:
                    if l <= abs(world[cx][cy] - world[nx][ny]) <= r:
                        visited[nx][ny] = united_num
                        q.append([nx, ny])
    while united:
        x, y = united.popleft()
        world[x][y] = (united_people // united_num)

    if united_num == 1:
        return 0
    return 1

while True:
    q = deque()
    break_cnt = 0
    visited = [[0 for _ in range(n)] for _ in range(n)]
    for i in range(n):
        for j in range(n):
            if visited[i][j] == 0:
                q.append((i, j))
                break_cnt += bfs(i, j)
    if break_cnt == 0:
        break
    else:
        day += 1

print(day)