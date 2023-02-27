import sys 
import heapq
input = sys.stdin.readline 
sys.setrecursionlimit(10**6)

v = int(input())
def max_dijkstra(s):
    cost = [float('inf') for _ in range(v+1)]
    q = []
    cost[0]=0
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

graph = [{}for _ in range(v+1)]
for _ in range(v):
    points = list(map(int, input().split()))
    s = points[0]
    for i in range(1, len(points)-1, 2):
        e, c = points[i], points[i+1]
        graph[s][e]=c
tmp = max_dijkstra(1)
idx = tmp.index(max(tmp))
tmp2 = max_dijkstra(idx)
print(max(tmp2))
