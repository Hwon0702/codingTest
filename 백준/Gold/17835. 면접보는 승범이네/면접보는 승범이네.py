import sys 
import heapq
input = sys.stdin.readline 
sys.setrecursionlimit(10**6)

def dijkstra():
    cost = [float('inf') for _ in range(n+1)]
    q = []
    for i in goals:
        heapq.heappush(q, [0, i])
        cost[i] = 0
    while q:
        current_cost, current_dest = heapq.heappop(q)
        if cost[current_dest]<current_cost:
            continue
        for next_dest, next_cost in graph[current_dest]:
            sum_cost = next_cost+current_cost
            if sum_cost<cost[next_dest]:
                cost[next_dest]=sum_cost
                heapq.heappush(q, [sum_cost, next_dest])
    return cost

n, m, k = map(int, input().split())
graph = [[]for _ in range(n+1)]

for _ in range(m):
    s, e, c = map(int, input().split())
    graph[e].append([s, c])
goals =  list(map(int, input().split()))
dist = dijkstra()
idx, max_dist = 0, 0
for i in range(1, n+1):
    if max_dist < dist[i]:
        idx, max_dist = i, dist[i]
print(idx)
print(max_dist)