import sys 
import math
import heapq
input = sys.stdin.readline
sys.setrecursionlimit(10**6)
n = int(input())
stars = [[] for _ in range(n)]
dir = [[] for _ in range(n)] 
for i in range(n):
    stars[i]=list(map(float, input().split()))
    dir[i] = [float('inf')]*(n)
#print(stars)
for i in range(n):
    for j in range(n):
        if i==j:
            continue
        else:
            dir[i][j]=math.sqrt(((stars[i][0]-stars[j][0]) ** 2) + ((stars[i][1]-stars[j][1]) ** 2))
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
        if current != next and not next in visit:
            heapq.heappush(q, (dir[current][next], next))
print("%.2f"% sum)