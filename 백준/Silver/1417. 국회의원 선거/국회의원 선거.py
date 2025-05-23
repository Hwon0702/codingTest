import sys 
import heapq
input = sys.stdin.readline
sys.setrecursionlimit(10**6)

hq = []
n = int(input())
res = -1*int(input())
cnt = 0
if n>1:
    for _ in range(n-1):
        heapq.heappush(hq, -1*int(input()))
    while True:
        p = heapq.heappop(hq)
        if (res < p):
            break
        else:
            cnt+=1
            res-=1
            heapq.heappush(hq, p+1)
print(cnt)
