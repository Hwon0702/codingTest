import sys
import heapq
input = sys.stdin.readline 
sys.setrecursionlimit(10**6)

n, m, k = map(int, input().split())

def dijkstra():
    cost = [[float('inf') for _ in range(k+1)] for _ in range(n)]
    q = []
    heapq.heappush(q, [0,0,0]) #start, count(K), dest
    cost[0][0] = 0
    while q:
        current_cost,current_cnt,current_dest = heapq.heappop(q)
        if current_cost > cost[current_dest][current_cnt]:
            continue
        if current_dest == n-1:
            return current_cost
        for next_dest,next_cost in graph[current_dest]:
            sum_cost = current_cost + next_cost
            if  sum_cost < cost[next_dest][current_cnt] :
                cost[next_dest][current_cnt] = sum_cost
                heapq.heappush(q,[sum_cost,current_cnt,next_dest])

            if current_cnt+1 <= k and cost[next_dest][current_cnt+1] > current_cost:
                cost[next_dest][current_cnt+1] = current_cost
                heapq.heappush(q, [current_cost, current_cnt+1, next_dest])

graph = [[] for _ in range(n+1)]

for _ in range(m):
    a, b, c = map(int, input().split())
    graph[a-1].append([b-1,c])
    graph[b-1].append([a-1,c])
print(dijkstra())