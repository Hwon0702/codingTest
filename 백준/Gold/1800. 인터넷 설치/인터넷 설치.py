import sys 
input = sys.stdin.readline 
sys.setrecursionlimit(10**6)
import heapq
#n이 목표, p쌍만 서로 연결, k는 공짜, 가장 비싼것만 가격을 받음 -> 가장 싸게 연결 필요
# 시작부터 n까지 가장 적은 비용으로 연결되게 
def dijkstra(now_payment):
    cost = [float('inf') for _ in range(n+1)]
    q = []
    cost[1]=0
    heapq.heappush(q, [0, 1])
    while q:
        current_cost,current_dest = heapq.heappop(q)
        if cost[current_dest] < current_cost: 
            continue
        for next_node, next_cost in lines[current_dest]:
            if now_payment < next_cost:
                if cost[next_node] > current_cost + 1:
                    cost[next_node] = current_cost + 1
                    heapq.heappush(q, [current_cost + 1, next_node])
            else:
                if cost[next_node] > current_cost:
                    cost[next_node] = current_cost
                    heapq.heappush(q, [current_cost, next_node])
    return cost[n]

n, p, k = map(int, input().split())
lines = [[] for _ in range(n+1)]
for _ in range(p):
    a, b, c = map(int, input().split())
    lines[a].append([b,c])
    lines[b].append([a,c])
res = float('inf')
left, right = 0, 1000000
while left <= right:
    mid = (left + right) // 2
    if dijkstra(mid) <= k:
        right = mid - 1
        res = mid
    else:
        left = mid + 1

if res == float('inf'): 
    print(-1)
else: 
    print(res)