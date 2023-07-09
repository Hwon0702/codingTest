import sys 
import heapq
input = sys.stdin.readline
sys.setrecursionlimit(10**6)

def dijkstra(start):
    cost = [float('inf') for _ in range(d+1)]
    q = []
    cost[start]=0
    heapq.heappush(q, [cost[start], start])
    while q:
        current_cost, current_dest = heapq.heappop(q)
        if cost[current_dest]<current_cost:
            continue
        for next_dest, next_cost in graph[current_dest]:
            sum_cost = next_cost+current_cost
            if sum_cost<cost[next_dest]:
                cost[next_dest]=sum_cost
                heapq.heappush(q, [sum_cost, next_dest])
    return cost[d]


n, d = map(int, input().split())
graph = [[]for _ in range(d+1)]

for i in range(d):
    graph[i].append([i+1, 1])

for _ in range(n):
    s, e, c = map(int, input().split())
    if e>d:
        continue
    graph[s].append([e, c])
print(dijkstra(0))