from collections import deque
import sys
input = sys.stdin.readline
sys.setrecursionlimit(10 ** 6)

dx = [-1,1,0,0]
dy = [0,0,-1,1]
q = deque()
n, m = map(int, input().split())
graph = []
visited = [[[[False] * m for _ in range(n)] for _ in range(m)] for _ in range(n)]

q = deque()
rx, ry, bx, by=0,0,0,0
for i in range(n):
    tmp = list(map(str, input().strip('\n')))
    for j in range(m):
        if tmp[j]=='R':
            tmp[j]='.'
            rx, ry = i, j
        if tmp[j]=='B':
            tmp[j]='.'
            bx, by = i, j
    graph.append(tmp)
q.append([rx, ry, bx, by, 1])

def move(x,y,dx,dy):
    cnt = 0
    while graph[x+dx][y+dy] != '#' and graph[x][y] != 'O':
        x += dx
        y += dy
        cnt +=1
    return x,y,cnt

def sol():
    while q:
        rx, ry, bx, by, cnt = q.popleft()
        if cnt > 10:
            return -1
        for i in range(4):
            nrx, nry, r_move = move(rx, ry, dx[i], dy[i])
            nbx, nby, b_move = move(bx, by, dx[i], dy[i])
            if graph[nbx][nby] == 'O':
                continue
            if graph[nrx][nry] == 'O':
                return cnt
            if nrx == nbx and nry == nby:
                if r_move > b_move:
                    nrx -= dx[i]
                    nry -= dy[i]
                else:
                    nbx -= dx[i]
                    nby -= dy[i]
            if visited[nrx][nry][nbx][nby] == False:
                visited[nrx][nry][nbx][nby] = True
                q.append([nrx, nry, nbx, nby, cnt + 1])
    return -1
print(sol())