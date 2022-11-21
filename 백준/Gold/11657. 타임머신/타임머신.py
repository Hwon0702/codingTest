import sys 
input = sys.stdin.readline
sys.setrecursionlimit(10**6)
n, m = map(int, input().split())
dist = [float('inf') for _ in range(n+1)]
graph = []
for _ in range(m):
    graph.append(list(map(int, input().split())))
def bellmanford():
    dist[1]=0
    for i in range(n+1):
        for current, next, cost in graph:
            if dist[current]!=float('inf') and dist[next]>dist[current]+cost:
                dist[next]=dist[current]+cost
                if i==n-1:
                    return True 
    return False

res = bellmanford()
if res:
    print(-1)
else:
    for i in range(2, n+1):
        if dist[i]!=float('inf'):
            print(dist[i])
        else:
            print(-1)