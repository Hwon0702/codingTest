import sys
input = sys.stdin.readline
def find(x, y, dir):
    global res
    if x == n - 1 and y == n - 1:
        res += 1
        return
    if dir == 0 or dir == 2: #가로
        if y + 1 < n and graph[x][y + 1] == 0:
            find(x, y + 1, 0)
    if dir == 1 or dir == 2:  #세로
        if x + 1 < n and graph[x + 1][y] == 0:
            find(x + 1, y, 1)
    if x + 1 < n and y + 1 < n:
        if graph[x + 1][y + 1] == 0 and graph[x + 1][y] == 0 and graph[x][y + 1] == 0:  # 대각선
            find(x + 1, y + 1, 2)

n = int(input())
graph = []
for _ in range(n):
    graph.append(list(map(int, input().split())))
res = 0
if graph[n - 1][n - 1] == 1:
    print(0)
else:
    find(0, 1, 0)
    print(res)