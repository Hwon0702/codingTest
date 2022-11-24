from collections import deque
import sys 
input = sys.stdin.readline
sys.setrecursionlimit(10**6)
r, c = map(int, input().split())
forest = []#[) for _ in range(r)]

dx = [-1,1,0,0]
dy = [0,0,-1,1]
q = deque()
water =deque()
visited = [[False for _ in range(c)] for _ in range(r)] 
for i in range(r):
    tmp = list(map(str, input().strip('\n')))
    for j in range(c):
        if tmp[j]=='S':
            q.append([i,j])
            continue
        if tmp[j]=='*':
            water.append([i, j])
    forest.append(tmp)

cnt = 0
while q:
    cnt+=1
    tmp_q=deque()
    while q:
        cx, cy = q.popleft()
        if forest[cx][cy]=='*':
            continue
        visited[cx][cy] = True

        for i in range(4):
            nx = cx+dx[i]
            ny = cy+dy[i]
            if 0<=nx<r and 0<=ny<c:
                if forest[nx][ny]=='D':
                    print(cnt)
                    sys.exit()
                elif forest[nx][ny]=='.' and visited[nx][ny]==False:
                    tmp_q.append([nx, ny])
                    visited[nx][ny]=True
    q = tmp_q
    tmp_q=deque()
    while water:
        cx, cy = water.popleft()
        for i in range(4):
            nx = cx + dx[i]
            ny = cy + dy[i]
            if 0 <= nx < r and 0 <= ny < c and forest[nx][ny]=='.':
                forest[nx][ny]='*'
                tmp_q.append([nx, ny])
    water = tmp_q
else:
    print("KAKTUS")