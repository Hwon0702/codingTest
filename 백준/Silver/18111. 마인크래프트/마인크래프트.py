import sys
input = sys.stdin.readline 
n, m, b = map(int, input().split())
graph = []
for _ in range(n):
    tmp = list(map(int, input().split()))
    graph.append(tmp)
res =  float('inf')
floor = 0

for target in range(257):
    max_target, min_target = 0, 0
    for i in range(n):
        for j in range(m):
            if graph[i][j] >= target:
                max_target += graph[i][j] - target
            else:
                min_target += target - graph[i][j]
    if max_target + b >= min_target:
        if min_target + (max_target * 2) <= res:
            res = min_target + (max_target * 2) 
            floor = target

print(res, floor)