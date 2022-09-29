import sys 
input = sys.stdin.readline
sys.setrecursionlimit(10**6)
n = int(input())
graph = []
for i in range(n):
    graph.append(list(map(int, input().split())))
res = [[0] * i for i in range(1, n+1)]
res[0][0] = graph[0][0]
for i in range(1, n):
    for j in range(i+1):
        if j == 0:
            res[i][j] = res[i-1][j] + graph[i][j]
        elif i == j:
            res[i][j] = res[i-1][j-1] + graph[i][j]
        else:
            res[i][j] = max(res[i - 1][j - 1], res[i - 1][j]) + graph[i][j]
print(max(res[-1]))