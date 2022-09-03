from collections import deque
import sys 
input = sys.stdin.readline
sys.setrecursionlimit(10**6)
x, y = map(int, input().split())
visited = [False for _ in range(100001)]
q = deque()
q.append((x, 0))
while q:
    n, t = q.popleft()
    if n==y:
        print(t)
        break
    if 0<=n*2<100001 and visited[n*2]==False:
        q.append((n*2, t))
        visited[n*2]=True
    if 0<=n-1<100001 and visited[n-1]==False:
        q.append((n-1, t+1))
        visited[n-1]=True
    if 0<=n+1<100001 and visited[n+1]==False:
        q.append((n+1, t+1))
        visited[n+1]=True
    