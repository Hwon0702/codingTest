import sys 
import heapq
input = sys.stdin.readline 
sys.setrecursionlimit(10**6)
n, m, r = map(int, input().split())
def dijkstra(start):
    q = []
    items_sum = 0
    visited = []
    heapq.heappush(q, [0, start])
    while q:
        current_dist, current_dest = heapq.heappop(q)
        if current_dest not in visited:
            if current_dist>m:
                continue
            items_sum+=items[current_dest]
            visited.append(current_dest)
            for new_dest, new_dist in graph[current_dest].items():
                heapq.heappush(q, [current_dist+new_dist, new_dest])
    return items_sum

graph = [{}for _ in range(n+1)]
items = list(map(int, input().split()))
res = 0
for _ in range(r):
    a, b, l = map(int, input().split())
    graph[a-1][b-1]=l
    graph[b-1][a-1]=l
for i in range(n):
    res = max(res, dijkstra(i))
print(res)