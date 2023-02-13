import sys 
from collections import deque
input = sys.stdin.readline 
sys.setrecursionlimit(10**6)

n, m = map(int, input().split())
dx = [-1, 1, 0, 0]
dy = [0, 0, -1, 1]
graph = []
q = deque()
coin1 = []
coin2 = []
def find():
    while q:
        coin1, coin2, cnt = q.popleft()
        if cnt>=10:
            return -1
        for i in range(4):
            nx1 = coin1[0]+dx[i]
            ny1 = coin1[1]+dy[i]
            nx2 = coin2[0]+dx[i]
            ny2 = coin2[1]+dy[i]
            if (0<=nx1<n and 0<=ny1<m) and (0<=nx2<n and 0<=ny2<m):
                if graph[nx1][ny1]=='#':
                    nx1 = coin1[0]
                    ny1 = coin1[1]
                if graph[nx2][ny2]=='#':
                    nx2 = coin2[0]
                    ny2 = coin2[1]
                q.append([[nx1,ny1], [nx2,ny2], cnt+1])
            elif not (0<=nx1<n and 0<=ny1<m) and (0<=nx2<n and 0<=ny2<m):
                return (cnt+1)
            elif (0<=nx1<n and 0<=ny1<m) and not (0<=nx2<n and 0<=ny2<m):
                return (cnt+1)

    return -1
for i in range(n):
    tmp =list(map(str, input().strip('\n')))
    for j in range(len(tmp)):
        if tmp[j]=='o':
            if len(coin1)<=0:
                coin1 = [i, j]
            else:
                coin2 = [i, j]
            tmp[j]='.'
    graph.append(tmp)
q.append([coin1, coin2, 0])
print(find())