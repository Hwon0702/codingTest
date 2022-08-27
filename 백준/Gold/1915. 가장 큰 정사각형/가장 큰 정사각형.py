import sys
input = sys.stdin.readline
sys.setrecursionlimit(10**6)

n, m = map(int, input().split())
graph = [list(map(int, list(input().strip()))) for _ in range(n)]
res = [[0]*m for _ in range(n)]
w = 0
for i in range(n):
    for j in range(m):
        if i == 0 or j == 0:
            res[i][j] = graph[i][j]
        elif graph[i][j] == 0:
            res[i][j] = 0
        else:
            res[i][j] = min(res[i - 1][j - 1], res[i - 1][j], res[i][j - 1]) + 1
        w = max(res[i][j], w)
print(w*w)