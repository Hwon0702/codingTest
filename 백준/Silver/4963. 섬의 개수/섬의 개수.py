from collections import deque
import sys
input = sys.stdin.readline
sys.setrecursionlimit(10**6)
from collections import deque
def bfs(x, y):
    dx = [1, -1, 1, -1, 1, -1, 0, 0]
    dy = [0, 0, 1, -1, -1, 1, 1, -1]
    graph[x][y] = 0
    q = deque()
    q.append([x, y])
    while(len(q)>0):
        x, y = q.popleft()
        for i in range(8):
            nx = x+dx[i]
            ny = y+dy[i]
            if 0<=nx<h and 0<=ny<w and graph[nx][ny]==1:
                graph[nx][ny]=0
                q.append([nx, ny])

while (True):
    w, h = map(int, input().split())
    if(w==0 and h==0):
        break
    graph = []
    count = 0
    for i in range(h):
        line = list(map(int, input().split()))
        graph.append(line)

    for i in range(h):
        for j in range(w):
            if graph[i][j]==1:
                bfs(i, j)
                count+=1
    print(count)
