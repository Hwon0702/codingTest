import sys
import heapq
input = sys.stdin.readline 
sys.setrecursionlimit(10**6)

n, m = map(int, input().split())

def dijkstra(start):
    cost = [float('inf') for _ in range(n+1)]
    cost[start]=0
    q = []
    heapq.heappush(q, [cost[start], start])
    while q:
        current_cost, current_dest = heapq.heappop(q)
        if cost[current_dest]<current_cost:
            continue
        for new_dest, new_cost in graph[current_dest]:
            sum_cost = current_cost+new_cost
            if cost[new_dest] and sum_cost<cost[new_dest]:
                cost[new_dest] = sum_cost
                heapq.heappush(q, [sum_cost, new_dest])
    return cost

graph = [[] for _ in range(n+1)]

for _ in range(m):
    a, b, c = map(int, input().split())
    graph[a].append([b,c])
    graph[b].append([a,c])
a = dijkstra(1)
print(a[n])