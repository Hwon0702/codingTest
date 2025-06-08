import sys 
sys.setrecursionlimit(10**6)

n, m = map(int, input().split())
graph = []
rect = 0

for _ in range(n):
    graph.append(list(map(int, input().strip())))

for i in range(n):
    for j in range(m):
        target = graph[i][j] 
        for k in range(j, m):
            if graph[i][k] == target and i + k - j < n and k < m:
                if graph[i + k - j][j] == target and graph[i + k - j][k] == target:
                    rect = max(rect, ((k - j + 1) ** 2))
print(rect)