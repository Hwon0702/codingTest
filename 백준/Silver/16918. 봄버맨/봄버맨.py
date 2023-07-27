import sys 
from collections import deque
input = sys.stdin.readline 
sys.setrecursionlimit(10**6)

dx=[-1,1,0,0]
dy=[0,0,-1,1]

r, c, n = map(int, input().split())
graph= []
q = deque()
for i in range(r):
    graph.append(list(map(str, input().rstrip('\n'))))
def bfs(q):
    while q:
        cx, cy = q.popleft()
        graph[cx][cy]='.'
        for i in range(4):
            nx = cx+dx[i]
            ny = cy+dy[i]
            if 0<=nx<r and 0<=ny<c and graph[nx][ny]=='O':
                graph[nx][ny]='.'


def bomb(t):
    global q, graph
    if t == 1:
        for i in range(r):
            for j in range(c):
                if graph[i][j] == 'O':
                    q.append([i,j])
    elif t%2 == 1:
        bfs(q)
        for i in range(r):
            for j in range(c):
                if graph[i][j] == 'O':
                    q.append([i,j])
    else:
        graph = [['O' for _ in range(c)] for _ in range(r)]

for i in range(1,n+1):
    bomb(i)

for i in graph:
    print(''.join(i))