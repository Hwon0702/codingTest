import sys 
input = sys.stdin.readline 
sys.setrecursionlimit(10**6)

def find(x):
    if parent[x] != x:
        parent[x] = find(parent[x])
    return parent[x]


def union(x, y):
    x, y = find(x), find(y)
    if x < y:
        parent[y] = x
    else:
        parent[x] = y

while True:
    m, n = map(int, input().split())
    if m == 0 and n == 0: 
        break
    parent = [i for i in range(m)]
    cost = 0
    graph = []
    for _ in range(n):
        x, y, z = map(int, input().split())
        graph.append((x, y, z))
    graph.sort(key=lambda x: x[2]) 

    for edge in graph:
        x, y, z = edge
        if find(x) != find(y): 
            union(x, y)
        else: 
            cost += z
    print(cost)