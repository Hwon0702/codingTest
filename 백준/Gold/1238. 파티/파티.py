import heapq
import sys
input = sys.stdin.readline
sys.setrecursionlimit(10**6)

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
N, M, X = map(int, input().split())
graph = {}
for i in range(1, M+1):
    graph[i]={}
for i in range(M):
    s, e, c = map(int, input().split())
    if graph[s] and e in graph[s] and graph[s][e]<c:
        continue
    else:
        graph[s][e]=c
max_cost = 0
for i in range(1, N+1):
    res = dijkstra(graph, i)
    res2 = dijkstra(graph, X)
    max_cost = max(max_cost, (res[X]+res2[i]))
print(max_cost)
#N개의 숫자로 구분된 각각의 마을에 한 명의 학생이 살고 있다.