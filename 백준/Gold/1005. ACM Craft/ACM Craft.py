from collections import deque
import sys
input = sys.stdin.readline
sys.setrecursionlimit(10**6)
def make_build():
    n, k = map(int, input().split())
    buildings = [[]for _ in range(n+1)]
    res_time = [0 for _ in range(n+1)]
    times = list(map(int, input().split()))
    times.insert(0, 0)
    cnt = [0 for _ in range(n+1)]
    q = deque()
    for _ in range(k):
        start, end = map(int, input().split())
        buildings[start].append(end)
        cnt[end]+=1
    w = int(input())
    for i in range(1, n+1):
        if cnt[i] == 0:
            q.append(i)
            res_time[i] += times[i]
    while q:
        now = q.popleft()
        for v in buildings[now]:
            cnt[v] -= 1
            res_time[v] = max(res_time[v], res_time[now] + times[v])
            if cnt[v] == 0:
                q.append(v)
    return res_time[w]

tc = int(input())
for _ in range(tc):
    print(make_build())