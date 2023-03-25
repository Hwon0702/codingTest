import sys 
input = sys.stdin.readline 
sys.setrecursionlimit(10**6)
n = int(input())
graph = []
for _ in range(n):
    tmp = list(map(str, input().split()))
    graph.append(tmp[1:])

graph.sort()

res = []
for i in range(n):
    if i == 0:
        for j in range(len(graph[i])):
            res.append('--' * j + graph[i][j])
    else:
        idx = 0 
        for j in range(len(graph[i])): #겹치는거 X : 새로운 루트
            if graph[i - 1][j] != graph[i][j] or len(graph[i - 1]) <= j:
                break
            else:#겹치는거 O : 브랜치
                idx = j + 1
        for j in range(idx, len(graph[i])):
            res.append('--' * j + graph[i][j])

for v in res:
    print(v)
