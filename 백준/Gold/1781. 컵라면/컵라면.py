import sys 
import heapq
input = sys.stdin.readline
sys.setrecursionlimit(10**6)
n = int(input())
lists = []
for _ in range(n):
    d, c = map(int, input().split())
    lists.append((d, c))
lists.sort()
q = []
for v in lists:
    heapq.heappush(q, v[1])
    if v[0]<len(q):
        heapq.heappop(q)
print(sum(q))