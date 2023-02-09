import sys 
input =sys.stdin.readline
sys.setrecursionlimit(10**6)
import heapq
def dijkstra(graph, s):
    cost = [float('inf') for _ in range(n+1)]
    q = []
    cost[s]=0
    heapq.heappush(q, [cost[s], s])
    while q:
        current_cost, current_dest = heapq.heappop(q)
        if cost[current_dest]<current_cost:
            continue
        for new_dest, new_cost in graph[current_dest].items():
            sum_cost = current_cost+new_cost
            if cost[new_dest] and sum_cost<cost[new_dest]:
                cost[new_dest] = sum_cost
                heapq.heappush(q, [sum_cost, new_dest])
    return cost
n, m, k, x = map(int, input().split())
graph = {}
for i in range(n+1):
    graph[i]={}
for _ in range(m):
    s, e = map(int, input().split())
    graph[s][e]=1
cost = dijkstra(graph, x)
res = []
for i in range(len(cost)):
    if cost[i]==k:
        res.append(i)
res.sort()
if len(res)<=0:
    print(-1)
else:
    for v in res:
        print(v)