import sys 
import heapq
input = sys.stdin.readline
sys.setrecursionlimit(10**6)
n=int(input())
days = {}
for i in range(n):
    d, c = map(int, input().split())
    if d not in days:
        days[d] = [c]
    else:
        days[d].append(c)
q = []
res = 0
remain = max(days.keys())
while 0 <= remain:
    if q:
        c = heapq.heappop(q)
        res += c*(-1)
    if remain in days:
        for x in days[remain]:
            heapq.heappush(q, -x)
    remain -= 1

print(res)