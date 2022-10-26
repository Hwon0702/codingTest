import heapq
import sys
input = sys.stdin.readline
sys.setrecursionlimit(10**6)
m, n = map(int, input().split())
graph = []
visited = [[False for _ in range(m)]for _ in range(n)]
cnt =[[0 for _ in range(m)]for _ in range(n)]
dx = [-1, 1, 0, 0]
dy = [0, 0, -1, 1]
for _ in range(n):
    graph.append(list(map(str, input().strip('\n'))))
q = []
heapq.heappush(q, [0, 0, 0])
while q:
    c, cx, cy = heapq.heappop(q)
    if cx ==n-1 and cy==m-1:
        break
    for i in range(4):
        nx = cx+dx[i]
        ny = cy+dy[i]
        if 0<=nx<n and 0<=ny<m and visited[nx][ny]==False:
            visited[nx][ny]=True 
            if graph[nx][ny]=='0':
                heapq.heappush(q, [c, nx, ny])
                cnt[nx][ny] = cnt[cx][cy]
            else:
                heapq.heappush(q, [c+1, nx, ny])
                cnt[nx][ny] = cnt[cx][cy]+1
print(cnt[n-1][m-1])