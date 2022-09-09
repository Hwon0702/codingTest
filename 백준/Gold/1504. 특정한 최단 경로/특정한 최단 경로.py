from cmath import inf
import sys 
input = sys.stdin.readline
sys.setrecursionlimit(10**6)
import heapq
def dijkstra(start, end):
    cost = {node: float('INF') for node in graph}  
    cost[start] = 0  
    queue = []
    heapq.heappush(queue, [cost[start], start])    
    while queue: 
      current_distance, current_destination = heapq.heappop(queue)      
      if cost[current_destination] < current_distance:  
        continue
      for new_destination, new_distance in graph[current_destination].items():
        distance = current_distance + new_distance 
        if distance < cost[new_destination]: 
          cost[new_destination] = distance
          heapq.heappush(queue, [distance, new_destination]) 
    return cost[end]

N, E = map(int, input().split())
graph = {}
for i in range(1, N+1):
    graph[i]={}
for _ in range(E):
    s, e, c = map(int, input().split())
    if graph[s] and e in graph[s] and graph[s][e]<c:
        continue
    else:
        graph[s][e]=c 
        graph[e][s]=c 
a, b = map(int, input().split())
res1 = dijkstra(1, a) + dijkstra(a, b) + dijkstra(b, N)
res2 = dijkstra(1, b) + dijkstra(b, a) + dijkstra(a, N)
if res1 ==float('INF')and res2==float('INF'):
    print(-1)
else:
    print(min(res1, res2))
