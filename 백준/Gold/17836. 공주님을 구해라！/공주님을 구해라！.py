from collections import deque
import sys 
input = sys.stdin.readline 
sys.setrecursionlimit(10**6)
n, m, t = map(int, input().split())
dx = [-1, 1, 0, 0]
dy = [0, 0, -1, 1]
sword = []
graph = []
sword_time = float('inf')
for i in range(n):
    tmp = list(map(int, input().split()))
    for j in range(m):
        if tmp[j]==2:
            s = [i, j]
    graph.append(tmp)
def find(start, end, k):
    global sword_time
    q = deque()
    q.append([start[0],start[1],0])
    visited = [[False for _ in range(m)]for _ in range(n)]
    while q:
        cx, cy, ti = q.popleft()
        if cx ==end[0] and cy ==end[1]:
            return  ti
        for i in range(4):
            nx = cx+dx[i]
            ny = cy+dy[i]
            if 0<=nx<n and 0<=ny<m and visited[nx][ny]==False:

                if not k:
                    if graph[nx][ny]==2:
                        sword_time = ti+1
                    if graph[nx][ny]!=1:
                        q.append([nx, ny, ti+1])
                        visited[nx][ny]=True
                else:
                    q.append([nx, ny, ti+1])
                    visited[nx][ny]=True
    return float('inf')
without_sword = find([0,0],[n-1, m-1], False)
with_sword = find(s,[n-1, m-1], True)
time = min(with_sword+sword_time, without_sword)
#print(sword_time, with_sword, without_sword, time)
if time<=t:
    print(time)
else:
    print("Fail")