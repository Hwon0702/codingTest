import sys 
input = sys.stdin.readline 
sys.setrecursionlimit(10**6)
import heapq
def dijkstra(graph, start):
    q = []
    heapq.heappush(q, [0, start])
    cost = [99999999 for _ in range(n+1)]
    cost[start] = 0
    while q:
        current_cost, current_dest= heapq.heappop(q)
        for new_dest, new_cost in graph[current_dest]:
            sum_cost = new_cost + current_cost
            if sum_cost < cost[new_dest]:
                cost[new_dest] = sum_cost
                heapq.heappush(q, [sum_cost, new_dest])
    return cost
tc = int(input())
for _ in range(tc):
    n, m, t = map(int, input().split())
    s, g, h = map(int, input().split())
    graph = [[] for _ in range(n + 1)]
    for i in range(m):
        a, b, d = map(int, input().split())
        graph[a].append([b, d])
        graph[b].append([a, d])
    goals = []
    res = []
    for i in range(t):
        goals.append(int(input()))
    start_to = dijkstra(graph, s)
    g_to = dijkstra(graph, g)
    h_to= dijkstra(graph, h)
    for i in goals:
        if start_to[g] + g_to[h] + h_to[i] == start_to[i] or start_to[h] + h_to[g] + g_to[i] == start_to[i]:
            res.append(i)
    
    res.sort()
    print(*res)