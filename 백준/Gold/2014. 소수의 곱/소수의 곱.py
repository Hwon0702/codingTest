import heapq
import sys 
input = sys.stdin.readline
sys.setrecursionlimit(10**6)
k, n = map(int, input().split())
primes = list(map(int, input().split()))
q = []
for v in primes:
    heapq.heappush(q, v)
for i in range(n):
    target = heapq.heappop(q)
    for j in primes:
        heapq.heappush(q, target*j)
        if target%j==0:
            break
print(target)