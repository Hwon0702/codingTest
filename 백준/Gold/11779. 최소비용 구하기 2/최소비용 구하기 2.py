import sys
input = sys.stdin.readline
sys.setrecursionlimit(10**6)
import heapq

def dijkstra(graph, s):
    cost = [float('inf') for _ in range(len(graph)+1)]
    road = [[]for _ in range(len(graph)+1)]
    q = []
    cost[s]=0
    road[s].append(s)
    print()
    heapq.heappush(q, [cost[s], s, road[s]])
    while q:
        current_cost, current_dest, current_road = heapq.heappop(q)
        if cost[current_dest]<current_cost:
            continue
        for new_dest, new_cost in graph[current_dest].items():
            sum_cost = current_cost+new_cost
            if sum_cost<cost[new_dest]:
                cost[new_dest] = sum_cost
                road[new_dest]=current_road+[new_dest]
                heapq.heappush(q, [sum_cost, new_dest, road[new_dest]])
    return cost, road

n = int(input())
m = int(input())
graph = {}
for i in range(1, n+1):
    graph[i]={}
for i in range(m):
    s, e, c = map(int, input().split())
    if graph[s] and e in graph[s] and graph[s][e]<c:
        continue
    else:
        graph[s][e]=c
start, end = map(int, input().split())
cost, road = dijkstra(graph, start)
print(cost[end])
print(len(road[end]))
for v in road[end]:
    print(v, end=' ')
