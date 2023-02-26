import sys 
from collections import deque
input = sys.stdin.readline
sys.setrecursionlimit(10**6)
dx = [-1, -1, -1, 0, 0, 0, 1, 1, 1]
dy = [-1, 0, 1, -1, 0, 1, -1, 0, 1]

n = int(input())
q = deque()
boards = []
bomb = [[0 for _ in range(n)]for _ in range(n)]
for i in range(n):
    tmp = list(map(str, input().rstrip('\n')))
    for j in range(len(tmp)):
        if tmp[j]=='#':
            q.append([i, j])
            bomb[i][j]=float('inf') #일단 가려진 부분은 폭탄이 있다고 가정하기
        else:
            tmp[j]=int(tmp[j])
    boards.append(tmp)
while q:
    cx, cy = q.popleft()
    for i in range(9):
        nx = cx+dx[i]
        ny = cy+dy[i]
        if 0<=nx<n and 0<=ny<n:
            if boards[nx][ny]!='#':
                bomb[cx][cy] = min(bomb[cx][cy], boards[nx][ny])
    if bomb[cx][cy]>0:
        for i in range(9):
            nx = cx+dx[i]
            ny = cy+dy[i]
            if 0<=nx<n and 0<=ny<n :
                if boards[nx][ny]!='#' and boards[nx][ny]>0:
                    boards[nx][ny] -=1
res = 0
for v in bomb:
    for k in v:
        if k>0:
            res+=1
print(res)