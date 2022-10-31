import sys 
from collections import deque
input = sys.stdin.readline
sys.setrecursionlimit(10**6)
def bfs():
    q = deque()
    q.append(S)
    visited[S]=True
    while q:
        now = q.popleft()
        if now ==G:
            return count[G]
        for v in (now-D, now+U):
            if 0<v<=F and visited[v]==False:
                visited[v]=True
                q.append(v)
                count[v]=count[now]+1
        
    if count[G]<=0:
        return "use the stairs"
F,S,G,U,D = map(int, input().split())
#총 F 목적 G 현재 S 
visited = [False for _ in range(F+1)]
count = [0 for _ in range(F+1)]
print(bfs())