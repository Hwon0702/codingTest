import sys 
input = sys.stdin.readline 
sys.setrecursionlimit(10**6)

h, w = map(int, input().split())
graph = [[] for _ in range(h)]
for i in range(h):
    graph[i] = list(map(str, input().strip('\n')))

dx = [-1, 0, 1]
dy = [1, 1, 1]

def move(x, y):
    global cnt 
    graph[x][y]='a'
    if y==w-1:
        cnt+=1
        return True
    for i in range(3):
        nx = x+dx[i]
        ny = y+dy[i]
        if 0<=nx<h and 0<=ny<w and graph[nx][ny]=='.':
            if move(nx, ny):
                return True
cnt = 0
for i in range(h):
    move(i,0)

print(cnt)