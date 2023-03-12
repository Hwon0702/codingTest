import sys 
input = sys.stdin.readline 
sys.setrecursionlimit(10**6)
dx = [0, -1, 0, 1]
dy = [1, 0, -1, 0]
n = int(input())
graph = [[0 for _ in range(101)] for _ in range(101)]
for i in range(n):
    y, x, d, g = map(int, input().split())
    graph[x][y] = 1
    gen = [d]
    q = [d]
    for _ in range(g + 1):
        for i in q:
            x += dx[i]
            y += dy[i]
            graph[x][y] = 1
        q = [(i + 1) % 4 for i in gen]
        q.reverse()
        gen = gen + q
result = 0
for i in range(100):
        for j in range(100):
            if graph[i][j] and graph[i][j + 1] and graph[i + 1][j] and graph[i + 1][j + 1]:
                result += 1
print(result)