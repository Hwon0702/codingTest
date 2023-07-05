import sys
import heapq 
input = sys.stdin.readline 
sys.setrecursionlimit(10**6)

n, m, d, e = map(int, input().split())
def dijkstra(start):
    cost = [float('inf') for _ in range(n+1)]
    q = []
    cost[start]=0
    heapq.heappush(q, [0, start])
    while q:
        current_cost,current_dest = heapq.heappop(q)
        if current_cost > cost[current_dest]:
            continue
        for next_dest,next_cost in graph[current_dest]:
            sum_cost = current_cost + next_cost
            if  sum_cost < cost[next_dest] and height[next_dest] > height[current_dest]:
                cost[next_dest] = sum_cost
                heapq.heappush(q,[sum_cost,next_dest])
    return cost

graph = [[]for _ in range(n+1)]
height = [0]+list(map(int, input().split()))
for _ in range(m):
    a, b, h = map(int, input().split())
    graph[a].append([b, h])
    graph[b].append([a, h])

home_to_school = dijkstra(1)
school_to_home = dijkstra(n)
res = []
for i in range(1, n+1):
    if home_to_school[i] != float('inf') and school_to_home[i] != float('inf'):
        res.append(height[i] * e - d * (home_to_school[i] + school_to_home[i]))
if len(res)>0:
    print(max(res))
else:
    print('Impossible')
