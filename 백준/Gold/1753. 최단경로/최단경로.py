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


p, l = map(int, input().split())
start = int(input())
graphs = {}
for i in range(1, p+1):
    graphs[i]={}
for i in range(l):
    s, e, c = map(int, input().split())
    if graphs[s] and e in graphs[s] and graphs[s][e]<c:
        continue
    else:
        graphs[s][e]=c
res = dijkstra(graphs, start)
for i in range(1, p+1):
    if i==start:
        print(0)
    elif res[i]==float('INF'):
        print("INF") 
    else:
        print(res[i])