import sys
input = sys.stdin.readline
sys.setrecursionlimit(10**6)
def dfs(x, y):
    dx = [1, -1, 0, 0] #상하좌우
    dy = [0, 0, 1, -1] #상하좌우

    for i in range(4):
        nx = x + dx[i]
        ny = y + dy[i]
        if (0 <= nx < n) and (0 <= ny < m): #범위내에 있을경우
            if graph[nx][ny] == 1: #연결되어있을경우
                graph[nx][ny] = 0  #연결돼있는데를 안가도 되는데로 만듦
                dfs(nx, ny)

tc = int(input().strip('\n'))
for _ in range(tc):
    m, n, k = map(int, input().split())
    graph = [[0 for _ in range(m)] for _ in range(n)]
    cnt = 0
    for _ in range(k):
        x, y = map(int, input().split())
        graph[y][x] = 1

    for i in range(n): 
        for j in range(m):  
            if graph[i][j]==1:
                dfs(i, j)
                cnt += 1
    print(cnt)