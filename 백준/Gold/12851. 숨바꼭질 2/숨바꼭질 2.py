from collections import deque
import sys
input = sys.stdin.readline
sys.setrecursionlimit(10**6)
def bfs(s, e):
    visited = [-1]*100001
    q = deque()
    q.append(s)
    visited[s]=0
    cnt=0
    while q:
        now = q.popleft()
        if now==e:
            cnt+=1
        if 0<=now*2<100001:
            if visited[now*2]==-1 or visited[now*2]==visited[now]+1:
                visited[now*2]=visited[now]+1
                q.append(now*2)
        if 0<=now+1<100001:
            if visited[now+1]==-1 or visited[now+1]==visited[now]+1:
                visited[now+1]=visited[now]+1
                q.append(now+1)
        if 0<=now-1<100001:
            if visited[now-1]==-1 or visited[now-1]==visited[now]+1:
                visited[now-1]=visited[now]+1
                q.append(now-1)
    print(visited[e])
    print(cnt)      
N, K = map(int, input().split())
bfs(N, K)