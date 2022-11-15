import sys
from collections import deque
input = sys.stdin.readline
sys.setrecursionlimit(10**6)
h, w = map(int, input().split())
graph = []
fx, fy = 0,0
dx=[-1,1,0,0]
dy=[0,0,-1,1]

m = deque()
f = deque()
def move():
    global m, f, graph
    visited = [[False for _ in range(w)]for _ in range(h)]
    move = []
    fire = []
    cnt=0
    while True:
        cnt+=1
        while f:
            cx, cy = f.popleft()
            for i in range(4):
                nx = cx+dx[i]
                ny = cy+dy[i]
                if 0<=nx<h and 0<=ny<w :
                    if graph[nx][ny]=='.':
                        fire.append([nx, ny])
                        graph[nx][ny]='F'
        f = deque(fire)
        fire=[]
        while m:
            cx, cy = m.popleft()
            if cx ==0 or cy==0 or cx==h-1 or cy==w-1:
                return cnt
            for i in range(4):
                nx = cx+dx[i]
                ny = cy+dy[i]
                if 0<=nx<h and 0<=ny<w and visited[nx][ny]==False and graph[nx][ny]=='.':
                    visited[nx][ny]=True
                    move.append([nx, ny])
        m = deque(move)
        move = []
        if not m :
            return "IMPOSSIBLE"
    return "IMPOSSIBLE"


for i in range(h):
    tmp = list(map(str, input().strip('\n')))
    for j in range(w):
        if tmp[j]=='J':
            m.append([i, j])
            tmp[j]='.'
        if tmp[j]=='F':
            f.append([i, j])
    graph.append(tmp)
print(move())