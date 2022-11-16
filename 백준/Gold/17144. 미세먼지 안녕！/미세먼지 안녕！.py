import sys
from collections import deque
input = sys.stdin.readline

h, w, t = map(int, input().split())
area = []
top, bottom = 0,0
for i in range(h):
    tmp = list(map(int, input().split()))
    for j in range(w):
        if top ==0 and bottom==0 and tmp[j]==-1:
            top = i
            bottom = i+1
    area.append(tmp)

def spread():
    dx = [-1, 1, 0, 0]
    dy = [0, 0, -1, 1]
    spread_area = [[0 for _ in range(w)]for _ in range(h)]
    for i in range(h):
        for j in range(w):
            if area[i][j]==0 or area[i][j]==-1:
                continue
            cx = i
            cy = j 
            spreaded = area[cx][cy]//5
            for k in range(4):
                nx = cx+dx[k]
                ny = cy+dy[k]
                if 0<=nx<h and 0<=ny<w and area[nx][ny]!=-1:
                    spread_area[nx][ny]+=spreaded 
                    spread_area[cx][cy]-=spreaded

    for i in range(h):
        for j in range(w):
            area[i][j]+=spread_area[i][j]
def top_clean():
    top_dx = [0, -1, 0, 1]
    top_dy=[1, 0, -1, 0]
    x, y, i = top, 1, 0
    prev = 0

    while True:
        nx = x + top_dx[i]
        ny = y + top_dy[i]
        if x == top and y == 0:
            break
        if not 0 <= nx < h or not 0 <= ny < w:
            i += 1
            continue

        area[x][y], prev = prev, area[x][y]
        x, y = nx, ny

def bottom_clean():
    bottom_dx=[0, 1, 0, -1]
    bottom_dy= [1, 0, -1, 0]
    x, y, i = bottom, 1, 0
    prev = 0

    while True:
        nx = x + bottom_dx[i] 
        ny = y + bottom_dy[i]
        if x == bottom and y == 0:
            break
        if not 0 <= nx < h or not 0 <= ny < w:
            i += 1
            continue

        area[x][y], prev = prev, area[x][y]
        x, y = nx, ny

for _ in range(t):
    spread()
    top_clean()
    bottom_clean()
res = 2
for v in area:
    res+=sum(v)
print(res)