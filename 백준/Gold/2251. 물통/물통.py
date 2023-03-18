from collections import deque
import sys 
input = sys.stdin.readline 
sys.setrecursionlimit(10**6)
water = deque()
origin_a, origin_b, origin_c = map(int, input().split())
visited = [[0] * 201 for i in range(201)]

water.append([0, 0, origin_c ])
res = set()
while water:
    a, b, c = water.popleft()
    if visited[a][b] : 
        continue
    visited[a][b] = True
    if a == 0: 
        res.add(c)
    if a + b > origin_b: 
        water.append([a + b - origin_b, origin_b, c])
    else: water.append([0, a + b, c])
    if a + c > origin_c: 
        water.append([a + c - origin_c, b, origin_c])
    else: water.append([0, b, a + c])
    if b + a > origin_a: 
        water.append([origin_a, b + a - origin_a, c])
    else: water.append([b + a, 0, c])
    if b + c > origin_c: 
        water.append([a, b + c - origin_c, origin_c])
    else: water.append([a, 0, b + c])
    if c + a > origin_a: 
        water.append([origin_a, b, c + a - origin_a])
    else: water.append([c + a, b, 0])
    if c + b > origin_b: 
        water.append([a, origin_b, c + b - origin_b])
    else: water.append([a, c + b, 0])
r = list(res)
r.sort()
print(*r)