import sys 
import heapq
input = sys.stdin.readline 
sys.setrecursionlimit(10**6)

def dijkstra(start, n):
    cost = [float('inf') for _ in range(n+1)]
    cost[start]=0
    q = []
    heapq.heappush(q, [0, start])
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

tc = int(input())
for _ in range(tc):
    n, d, c = map(int, input().split())
    graph = [[]for _ in range(n+1)]
    for _ in range(d):
        a, b, s = map(int, input().split())
        graph[b].append([a, s])
    cost = dijkstra(c, n)
    max_v = 0 
    cnt = 0
    for v in cost:
        if v!=float('inf'):
            max_v = max(v, max_v)
            cnt+=1
    print(cnt, max_v)
    
