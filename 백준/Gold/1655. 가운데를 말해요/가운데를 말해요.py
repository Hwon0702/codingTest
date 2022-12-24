import sys 
input = sys.stdin.readline 
sys.setrecursionlimit(10**6)
import heapq
lq =[]
rq =[]
num = int(input())
res = []
for i in range(num):
    n = int(input())
    if len(lq)==len(rq):
        heapq.heappush(lq, (-n, n))
    else:
        heapq.heappush(rq, (n, n))
    if rq and lq[0][1]>rq[0][0]:
        min = heapq.heappop(rq)[0]
        max = heapq.heappop(lq)[1]
        heapq.heappush(lq, (-min, min))
        heapq.heappush(rq, (max, max))
    res.append(lq[0][1])
for n in res:
    print(n)
