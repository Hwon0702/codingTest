import sys 
import heapq
input = sys.stdin.readline 
sys.setrecursionlimit(10**6)

def dijkstra(start):
    cost = [[float('inf') for _ in range(k)] for _ in range(n + 1)]
    q = []
    cost[start][0]=0
    heapq.heappush(q, [0, start])
    while q:
        current_cost, current_dest = heapq.heappop(q)
        for next_dest, next_cost in graph[current_dest]:
            sum_cost = next_cost+current_cost
            if sum_cost < cost[next_dest][k-1]:
                cost[next_dest][k-1]=sum_cost
                cost[next_dest].sort()
                heapq.heappush(q, [sum_cost, next_dest])
    return cost

n, m, k = map(int, input().split())
graph = [[]for _ in range(n+1)]
for _ in range(m):
    s, e, c = map(int, input().split())
    graph[s].append([e, c])
cost = dijkstra(1)
for i in range(1, n + 1):
    if cost[i][k-1] == float('inf'):
        print(-1)
    else:
        print(cost[i][k-1])