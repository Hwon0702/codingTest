import sys
input = sys.stdin.readline
sys.setrecursionlimit(10 ** 6)


def dfs(x, y):
    if dp[x][y]: 
        return dp[x][y] # 이미 다녀온 곳
    dp[x][y] = 1
    for i in range(4):
        nx = x + dx[i]
        ny = y + dy[i]
        if 0 <= nx < n and 0 <= ny < n and graph[x][y] < graph[nx][ny]:
            dp[x][y] = max(dp[x][y], dfs(nx, ny) + 1)
    return dp[x][y]


n = int(input())
graph = []
for _ in range(n):
    graph.append(list(map(int, input().split())))
dx = [-1, 1, 0, 0]
dy = [0, 0, -1, 1]
dp = [[0] * n for _ in range(n)]

ans = 0
for i in range(n):
    for j in range(n):
        ans = max(ans, dfs(i, j))

print(ans)