import sys
from collections import deque

n, k = map(int, sys.stdin.readline().split())
visited = [0 for _ in range(100001)]
queue = deque()
queue.append(n)
while(queue):
    now = queue.popleft()
    if now == k:
        print(visited[now])
    else:
        for after in (now-1, now+1, now*2):
            if 0 <= after <= 100000 and not visited[after]:
                visited[after] = visited[now] + 1
                queue.append(after)