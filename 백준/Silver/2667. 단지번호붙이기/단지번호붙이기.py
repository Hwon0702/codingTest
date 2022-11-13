import sys 
input = sys.stdin.readline
sys.setrecursionlimit(10**6)
from collections import deque

def bfs(graph, x, y):
    dx = [-1,1,0,0]
    dy = [0,0,-1,1]
    queue = deque()
    queue.append((x,y))
    graph[x][y] = '0' 
    cnt = 1   
    while queue:
        x,y = queue.popleft()
        for i in range(4):
            nx = x + dx[i]
            ny = y + dy[i]
            if nx < 0 or nx >= n or ny <0 or ny >= n:
                continue
            if graph[nx][ny] == '1':
                graph[nx][ny] = 0
                queue.append((nx,ny))
                cnt += 1
    return cnt
n = int(input())
graph=[]
for i in range(n):
    graph.append(list(map(str, input().strip('\n'))))
count = []
for i in range(n):
    for j in range(n):
        if graph[i][j]=='1':
            count.append(bfs(graph, i, j))
count.sort()
print(len(count))
for i in range(len(count)):
    print(count[i])

