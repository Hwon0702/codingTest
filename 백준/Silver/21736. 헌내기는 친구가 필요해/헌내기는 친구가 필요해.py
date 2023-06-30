import sys 
from collections import deque
input = sys.stdin.readline 
sys.setrecursionlimit(10**6)
dx = [-1,1,0,0]
dy = [0,0,-1,1]
def search(x,y):
    q = deque()
    q.append([x,y])  
    
    dx = [1,-1,0,0]
    dy = [0,0,-1,1]
    
    visited = [[False for _ in range(m)] for _ in range(n)]
    
    cnt = 0
    while q:
        cx,cy = q.popleft()
        
        for i in range(4):
            nx = cx+dx[i]
            ny = cy+dy[i]
            
            if 0 <= nx < n and 0 <= ny < m and visited[nx][ny]==False and graph[nx][ny] != 'X':
                visited[nx][ny] = True
                if graph[nx][ny] == 'P':
                    cnt += 1
                q.append([nx,ny])
                    
    return cnt
      
n,m = map(int,input().split())
graph = []
sx, sy = -1,-1
for i in range(n):
    tmp = list(map(str, input().rstrip('\n')))
    for j in range(m):
        if tmp[j]=='I':
            sx, sy = i, j
    graph.append(tmp)
v = search(sx, sy)
if v==0:
    print('TT')
else:
    print(v)