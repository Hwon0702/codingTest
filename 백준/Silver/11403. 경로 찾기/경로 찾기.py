import sys 
input = sys.stdin.readline
sys.setrecursionlimit(10**6)
n = int(input())

graph = [list(map(int, input().split())) for _ in range(n)]
for k in range(0, n):
    for a in range(0, n ):
        for b in range(0, n ):
            if graph[a][k] and graph[k][b]:
                graph[a][b]=1
for v in graph:
    for n in v:
        print(n, end=' ')
    print()
