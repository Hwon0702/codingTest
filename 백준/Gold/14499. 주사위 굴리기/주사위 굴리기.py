import sys 
input = sys.stdin.readline 
sys.setrecursionlimit(10**6)

n, m, x, y, k = map(int, input().split())

graph = []
dx = [0,0,-1,1]
dy = [1,-1,0,0]
dice = [0, 0, 0, 0, 0, 0]

def turn(dir):
    a, b, c, d, e, f = dice[0], dice[1], dice[2], dice[3], dice[4], dice[5]
    if dir == 1: #동
        dice[0], dice[1], dice[2], dice[3], dice[4], dice[5] = d, b, a, f, e, c

    elif dir == 2: #서 
        dice[0], dice[1], dice[2], dice[3], dice[4], dice[5] = c, b, f, a, e, d
    elif dir == 4: #남
        dice[0], dice[1], dice[2], dice[3], dice[4], dice[5] = b, f, c, d, a, e
    else: #북
        dice[0], dice[1], dice[2], dice[3], dice[4], dice[5] = e, a, c, d, f, b

for i in range(n):
    graph.append(list(map(int, input().split())))

methods = list(map(int, input().split()))

nx, ny = x, y
for method in methods:
    nx += dx[method-1]
    ny += dy[method-1]
    if 0<=nx<n and 0<=ny<m:
        turn(method)
        if graph[nx][ny] == 0:
            graph[nx][ny] = dice[-1]
        else:
            dice[-1] = graph[nx][ny]
            graph[nx][ny] = 0

        print(dice[0])
    else:
        nx -= dx[method-1]
        ny -= dy[method-1]
        continue