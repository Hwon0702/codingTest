import sys 
input = sys.stdin.readline
sys.setrecursionlimit(10**6)
n = int(sys.stdin.readline())
graph = [list(map(int, input().split())) for _ in range(n)]

def dfs(x, y, z):
    global res
    visited = graph[x][y]
    for i in range(x, x + z):
        for j in range(y, y + z):
            if graph[i][j] != visited:
                for k in range(3):
                    for l in range(3):
                        dfs(x + k * z // 3, y + l * z // 3, z // 3)
                return
    if visited == -1:
        res[0] += 1
    elif visited == 0:
        res[1] += 1
    else:
        res[2] += 1

res = [0, 0, 0]
dfs(0, 0, n)
for v in res:
    print(v)
