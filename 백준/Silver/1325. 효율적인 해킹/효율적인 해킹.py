import sys 
from collections import deque
input = sys.stdin.readline 

def bfs(start):
	connected = 1
	q = deque()
	q.append(start)
	visit = [False for _ in range(n+1)]
	visit[start] = True

	while q:
		current = q.popleft()
		for next in computers[current]:
			if not visit[next]:
				visit[next] = True
				connected += 1
				q.append(next)

	return connected

n, m = map(int, input().split())
computers = [[] for _ in range(n+1)]
for _ in range(m):
    a, b = map(int, input().split())
    computers[b].append(a)

max_cnt = 0
tmp = []
for i in range(1, n+1):
	cnt = bfs(i)
	if cnt>max_cnt:
		max_cnt = cnt 
		tmp = [i]
	elif cnt==max_cnt:
		tmp.append(i)
print(*tmp)