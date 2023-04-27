import sys;
input = sys.stdin.readline
sys.setrecursionlimit(10**6)
dx = [-1, 1, 0, 0]
dy = [0, 0, 1, -1]
def dfs(r, c, idx, total):
    global res
    if res >= total + max_val * (3 - idx):
        return
    if idx == 3:
        res = max(res, total)
        return
    else:
        for i in range(4):
            nx = r + dx[i]
            ny = c + dy[i]
            if 0 <= nx < n and 0 <= ny < m and visited[nx][ny] == 0:
                if idx == 1:
                    visited[nx][ny] = 1
                    dfs(r, c, idx + 1, total + arr[nx][ny])
                    visited[nx][ny] = 0
                visited[nx][ny] = 1
                dfs(nx, ny, idx + 1, total + arr[nx][ny])
                visited[nx][ny] = 0


n, m = map(int, input().split())
arr = [list(map(int, input().split())) for _ in range(n)]
visited = [[False for _ in range(m)] for _ in range(n)]
res = 0
max_val = max(map(max, arr))
for r in range(n):
    for c in range(m):
        visited[r][c] = 1
        dfs(r, c, 0, arr[r][c])
        visited[r][c] = 0

print(res)
