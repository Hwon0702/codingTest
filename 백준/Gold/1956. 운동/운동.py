import sys 
input = sys.stdin.readline
n, m = map(int,input().split())
roads = [[float('inf')] * (n+1) for _ in range(n+1)]
for _ in range(m):
    s, e, c = map(int, input().split())
    roads[s][e] = min(roads[s][e], c)

for k in range(1, n + 1):
    for a in range(1, n + 1):
        for b in range(1, n + 1):
            roads[a][b] = min(roads[a][b], roads[a][k] + roads[k][b])
ans = float('inf')
for i in range(1, n + 1):
    ans = min(ans, roads[i][i])
if ans==float('inf'):
    print(-1)
else:
    print(ans)