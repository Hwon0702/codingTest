import sys 
import heapq
input = sys.stdin.readline
sys.setrecursionlimit(10**6)

def dijkstra(start):
    cost = [float('inf') for _ in range(n+1)]
    q = []
    cost[start]=0
    res_route = [0 for _ in range(n+1)]
    heapq.heappush(q, [cost[start], start])
    while q:
        current_cost, current_dest = heapq.heappop(q)
        if cost[current_dest]<current_cost:
            continue
        for next_dest, next_cost in graph[current_dest]:
            sum_cost = next_cost+current_cost
            if sum_cost<cost[next_dest]:
                cost[next_dest]=sum_cost
                res_route[next_dest]=current_dest
                heapq.heappush(q, [sum_cost, next_dest])
    return cost, res_route


n, m = map(int, input().split())
graph = [[]for _ in range(n+1)]
for _ in range(m):
    s, e, c = map(int, input().split())
    graph[s].append((e, c))
    graph[e].append((s, c))
res_cost, route = dijkstra(1)
print(n-1)
for i in range(2, n + 1):
    print(i, route[i])
