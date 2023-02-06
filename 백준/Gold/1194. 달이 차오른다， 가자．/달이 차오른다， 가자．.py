import sys 
from collections import deque
input = sys.stdin.readline 
sys.setrecursionlimit(10**6)
dx = [-1, 1, 0, 0]
dy = [0, 0, -1, 1]
n, m = map(int, input().split())

visited =[[[0] * 64 for i in range(m)] for i in range(n)]
graph = []
q = deque()
for i in range(n):
    tmp = list(map(str, input().strip('\n')))
    for j in range(len(tmp)):
        if tmp[j]=='0':
            q.append([i, j, 0, 0])
            visited[i][j][0]=0
            tmp[j]='.'
    graph.append(tmp)

def find():
    while q:
        x, y, key, cnt = q.popleft()
        for i in range(4):
            nx = x + dx[i]
            ny = y + dy[i]
            if 0 <= nx < n and 0 <= ny < m and graph[nx][ny] != "#" and visited[nx][ny][key]==0:
                if graph[nx][ny] == ".":
                    visited[nx][ny][key] = 1
                    q.append([nx, ny, key, cnt + 1])
                elif graph[nx][ny] == "1":
                    return cnt + 1
                else:
                    if graph[nx][ny].isupper():
                        if key & 1 << (ord(graph[nx][ny]) - 65):
                            visited[nx][ny][key] = 1
                            q.append([nx, ny, key, cnt + 1])
                    elif graph[nx][ny].islower():
                        nc = key | (1 << ord(graph[nx][ny]) - 97)
                        if visited[nx][ny][nc] == 0:
                            visited[nx][ny][nc] = 1
                            q.append([nx, ny, nc, cnt + 1])
    return -1

print(find())


