import sys 
import heapq
from heapq import merge
input = sys.stdin.readline
n, k = map(int, input().split())
bags = []
jewels = []
able_jewels=[]
max_cost = 0
for idx in range(n):
    weight, cost = map(int, input().split())
    heapq.heappush(jewels, (weight, cost))
for _ in range(k):
    bags.append(int(input()))
bags.sort()
for bag in bags:
    while jewels and bag>=jewels[0][0]:
        weight, cost = heapq.heappop(jewels)
        heapq.heappush(able_jewels, -cost)
    if able_jewels:
        able = heapq.heappop(able_jewels)
        max_cost+=able
    elif not jewels:
        break
print(-max_cost)
