import sys 
import heapq
input = sys.stdin.readline
sys.setrecursionlimit(10**6)

n = int(input())
m = int(input())
graph = [[10001 for _ in range(n+1)] for _ in range(n+1)]
visited = [False for _ in range(n+1)]

for _ in range(m):
    s, e, c = map(int, input().split())
    graph[s][e]=min(graph[s][e],c) 
    graph[e][s]=min(graph[e][s],c) 

res = 0
h = [[0, 1]]
while h:
    c, s = heapq.heappop(h)
    if visited[s]==False:
        visited[s]=True 
        res+=c
        for idx in range(1, n+1):
            if visited[idx]==False and graph[s][idx]<10001:
                heapq.heappush(h, [graph[s][idx], idx])
print(res)
