import sys
input = sys.stdin.readline
from collections import deque
from itertools import combinations
sys.setrecursionlimit(10**6)

def bfs(q):
    visited = [[False] * n for _ in range(n)]
    for x, y in q:
        visited[x][y] = True
    time = -1
    cnt = len(q)
    while q:
        for _ in range(len(q)):
            x, y = q.popleft()
            for i in range(4):
                nx, ny = x + dx[i], y + dy[i]
                if 0 <= nx < n and 0 <= ny < n and not visited[nx][ny] and not graph[nx][ny] == 1:
                    visited[nx][ny] = True
                    q.append((nx, ny))
                    cnt += 1
        time += 1
    return cnt, time

dx = [0, 0, -1, 1]
dy = [1, -1, 0, 0]

n, m = map(int, input().split())
graph = []
virus = []
wall = 0
for i in range(n):
    a = list(map(int, input().split()))
    for j in range(n):
        if a[j] == 2:
            virus.append([i,j])
        elif a[j] == 1:
            wall += 1
    graph.append(a)

min_count = float('inf')
for virus_list in combinations(virus, m):
    q = deque()
    for v in virus_list:
        q.append(v)
    cnt, time = bfs(q)
    if cnt + wall == n**2:
        min_count = min(min_count, time)
if min_count == float('inf'):
    min_count = -1
print(min_count)