import sys 
input = sys.stdin.readline
sys.setrecursionlimit(10**6)
def find(x):
    if parent[x] != x:
        parent[x] = find(parent[x])
    return parent[x]

def union(a, b):
    a = find(a)
    b = find(b)
    if a < b:
        parent[b] = a
    else:
        parent[a] = b

n = int(input())
parent = list(range(n))
point_x = []
point_y = []
point_z = []

for idx in range(n):
    a, b, c = map(int, input().split())
    point_x.append([a, idx])
    point_y.append([b, idx])
    point_z.append([c, idx])
point_x.sort()
point_y.sort()
point_z.sort()
dists = []
for i in range(n - 1):
    dists.append((point_x[i + 1][0] - point_x[i][0], point_x[i][1], point_x[i + 1][1]))
    dists.append((point_y[i + 1][0] - point_y[i][0], point_y[i][1], point_y[i + 1][1]))
    dists.append((point_z[i + 1][0] - point_z[i][0], point_z[i][1], point_z[i + 1][1]))

dists.sort()
answer = 0

for e in dists:
    cost, a, b = e
    if find(a) != find(b):
        union(a, b)
        answer += cost
print(answer)