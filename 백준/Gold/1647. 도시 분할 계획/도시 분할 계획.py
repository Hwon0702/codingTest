import sys , heapq
input = sys.stdin.readline 
n, m = map(int,input().split())
lines = [[]for _ in range(n)]
visited = [False]*n
visited_cnt = n
hq = []
max_cost, min_cost = 0, 0 
for _ in range(m):
    s, e, c = map(int, input().split())
    lines[s-1].append([e-1,c])
    lines[e-1].append([s-1,c])
heapq.heappush(hq, (0,0))
while visited_cnt>0 and hq:
    cost, current = heapq.heappop(hq)
    if visited[current]:
        continue
    visited[current]=True
    visited_cnt-=1
    max_cost = max(cost, max_cost)
    min_cost+=cost
    for _, v in enumerate(lines[current]):
        idx = v[0]
        c = v[1]
        if current!=idx and not visited[idx] and v!=float('inf'):
            heapq.heappush(hq, (c,idx))
print(min_cost-max_cost)