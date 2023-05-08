import sys
input = sys.stdin.readline 
sys.setrecursionlimit(10**6)
def floyd_warshal():
    result = [[j for j in range(n+1)] for i in range(n+1)]
    for k in range(1, n+1):
        result[k][k]=0
        for i in range(1, n+1):
            for j in range(1, n+1):
                if graph[i][j] > graph[i][k]+graph[k][j]:
                    graph[i][j] = graph[i][k]+graph[k][j]
                    result[i][j] = result[i][k]
    return result
n, m = map(int, input().split())
graph = [[float('inf') for _ in range(n+1)] for _ in range(n+1)]
for i in range(1, n+1):
    graph[i][i] = 0
for i in range(m):
    s, e, c = map(int, input().split())
    graph[s][e]=min(graph[s][e], c)
    graph[e][s]=min(graph[e][s], c)
res = floyd_warshal()
row = len(res)
col = len(res[0])
for i in range(1, row):
    for j in range(1, col):
        if i==j:
            print('-', end=' ')
        else:
            print(res[i][j], end=' ')
    print()