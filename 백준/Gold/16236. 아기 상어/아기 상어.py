from collections import deque
import sys 
input = sys.stdin.readline
sys.setrecursionlimit(10**6)
n = int(input())
graph = []
sx, sy = 0,0
now_size=2
dx = [0,0,-1,1]
dy = [-1,1,0,0]
min_cnt=999999
time = 0
def find(sx, sy):
    #print("SIZE",  sx, sy, now_size)
    visited= [[False for _ in range(n)]for _ in range(n)]
    q = deque()
    q.append([sx, sy, 0])
    visited[sx][sy]=True
    targets = []
    while q:
        cx, cy, cnt = q.popleft()
        if graph[cx][cy]>0 and graph[cx][cy]<now_size:
            targets.append([cnt, cx, cy])
        for i in range(4):
            nx = cx+dx[i]
            ny = cy+dy[i]
            if 0<=nx<n and 0<=ny<n and now_size>=graph[nx][ny] and visited[nx][ny]==False:
                visited[nx][ny]=True
                q.append([nx, ny, cnt+1])
    targets.sort()
    return targets




for i in range(n):
    tmp=list(map(int, input().split()))
    for j in range(n):
        if tmp[j]==9:
            sx, sy = i, j
            tmp[j]=0
    graph.append(tmp)
break_flag=False
while min_cnt>0 and  not break_flag:
    for _ in range(now_size):
        graph[sx][sy]=0
        res = find(sx, sy)
        if len(res)<=0:
            break_flag=True
            break
        sx = res[0][1]
        sy = res[0][2]
        min_cnt = res[0][0]
        graph[sx][sy]=0
        time+=min_cnt

    now_size+=1
print(time)