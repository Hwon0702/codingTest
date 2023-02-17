import sys 
import heapq
input = sys.stdin.readline 
sys.setrecursionlimit(10**6)
n = int(input())
num = []
for _ in range(n):
    heapq.heappush(num, -int(input()))
while num:
    print(-heapq.heappop(num))