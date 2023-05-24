import sys 
input = sys.stdin.readline 
sys.setrecursionlimit(10**6)
from itertools import combinations
from collections import deque
dx = [-1,1,0,0]
dy = [0,0,-1,1]

def find(start_list):
    q = deque()
    visited = [[-1 for _ in range(n)]for _ in range(n)]
    cnt = 0
    for v in start_list:
        q.append(v)
        visited[v[0]][v[1]]=0
    while q:
        cx, cy = q.popleft()
        for i in range(4):
            nx = cx+dx[i]
            ny = cy+dy[i]
            if 0<=nx<n and 0<=ny<n and visited[nx][ny]==-1 and graph[nx][ny]!=1:
                visited[nx][ny] = visited[cx][cy]+1
                q.append([nx, ny])
                cnt = max(cnt, visited[nx][ny])
    for row in range(n):
        for col in range(n):
            if visited[row][col] ==-1 and graph[row][col] ==0:
                return -1
    return cnt


    #q.append(start_list)



n, m = map(int, input().split())
graph = []
virus = []
for i in range(n):
    tmp = list(map(int, input().split()))
    for j in range(len(tmp)):
        if tmp[j]==2:
            virus.append([i, j])
    graph.append(tmp)
combination_list = list(combinations(virus, m))
res =float('inf')
min_v = -1
for comb in combination_list:
    a = find(comb)
    if a==-1 and res!=-1:
        continue
    elif a==-1 and res ==float('inf'):
        res = -1 
    elif a!=-1 and res==-1:
        res = a
    else:
        res = min(res, a)
if res==float('inf'):
    res = -1
print(res)