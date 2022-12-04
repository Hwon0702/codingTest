import heapq
import sys 
input = sys.stdin.readline
sys.setrecursionlimit(10**6)
n,m=map(int, input().split())
univ = list(map(str, input().split()))

lists = [[float('inf') for _ in range(n)]for _ in range(n)]
for _ in range(m):
    s, e, c = map(int, input().split())
    if univ[s-1]!=univ[e-1]:
        lists[s-1][e-1] = min(c , lists[s-1][e-1])
        lists[e-1][s-1] = min(c , lists[e-1][s-1])

sum = 0
visit = set()
visit_cnt = n
q = []
heapq.heappush(q, (0, 0))
while visit_cnt>0 and q:
    cost, current = heapq.heappop(q)
    if current in visit:
        continue
    visit.add(current)
    visit_cnt-=1
    sum += cost
    for next in range(n):
        if current != next and not next in visit and next!=float('inf'):
            heapq.heappush(q, (lists[current][next], next))
if visit_cnt>0 or sum==float('inf'):
    print(-1)
else:
    print(sum)