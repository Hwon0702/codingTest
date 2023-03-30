import sys
from collections import deque
input = sys.stdin.readline
sys.setrecursionlimit(10**6)
n, k = map(int,input().split())
visited=[[False]*500001 for _ in range(2)]
visited[0][n] = True
res = -1

q = deque()
q.append(n)
flag = 0
depth = 0
while q:
    if k > 500000: 
        break
    if visited[flag][k]: 
        res = depth
        break

    tmp =deque()
    flag = 1 - flag
    for x in q:
        if 0 <= x-1 <= 500000 and not visited[flag][x-1]:
            visited[flag][x-1] = True
            tmp.append(x-1)
        if 0 <= x+1 <= 500000 and not visited[flag][x+1]:
            visited[flag][x+1] = True
            tmp.append(x+1)
        if 0 <= x*2 <= 500000 and not visited[flag][x*2]:
            visited[flag][x*2] = True
            tmp.append(x*2)

    depth += 1
    k += depth
    q = tmp

print(res)