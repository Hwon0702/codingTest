import sys
input = sys.stdin.readline
sys.setrecursionlimit(10**6)
import heapq  # 우선순위 큐 구현을 위함

def dijkstra(graph, start):
  distances = {node: float('INF') for node in graph}  
  distances[start] = 0  
  queue = []
  heapq.heappush(queue, [distances[start], start])  

  while queue: 
    current_distance, current_destination = heapq.heappop(queue)  

    if distances[current_destination] < current_distance:  
      continue
    
    for new_destination, new_distance in graph[current_destination].items():
      distance = current_distance + new_distance 
      if distance < distances[new_destination]: 
        distances[new_destination] = distance
        heapq.heappush(queue, [distance, new_destination])  
    
  return distances

n = int(input())
m = int(input())
graphs = {}
for i in range(1, n+1):
    graphs[i]={}
for i in range(m):
    s, e, c = map(int, input().split())
    if graphs[s] and e in graphs[s] and graphs[s][e]<c:
        continue
    else:
        graphs[s][e]=c
start, end = map(int, input().split())
res = dijkstra(graphs, start)
print(res[end])