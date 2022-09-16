import heapq
import sys
input = sys.stdin.readline
sys.setrecursionlimit(10**6)
v, e = map(int, input().split())
lines = [[] for _ in range(v)]
visited = [False]*(v)
visit_count = v
sum = 0
q = []
for _ in range(e):
    s, d, c = map(int, input().split())
    lines[s-1].append((d-1, c))
    lines[d-1].append((s-1, c))
heapq.heappush(q, (0,0))
while visit_count >0:
    cost, current = heapq.heappop(q)
    if not visited[current]:
        visited[current] = True
        sum += cost
        visit_count -=1
        for next in lines[current]:
            if not visited[next[0]]:
                heapq.heappush(q, (next[1], next[0]))
print(sum)