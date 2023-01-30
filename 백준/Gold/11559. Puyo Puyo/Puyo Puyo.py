import sys 
from collections import deque
input = sys.stdin.readline 
sys.setrecursionlimit(10**6)
fields = []
q = deque()
dx = [-1, 1, 0, 0]
dy = [0, 0, -1, 1]


def find(i, j):
    q = deque()
    q.append([i, j])
    removed = []
    removed.append([i, j])
    global visited
    while q:
        cx, cy = q.popleft()
        for i in range(4):
            nx = cx+dx[i]
            ny = cy+dy[i]
            if 0<=nx<12 and 0<=ny<6 and fields[nx][ny]==fields[cx][cy] and visited[nx][ny]==False:
                    visited[nx][ny]=True 
                    q.append([nx,ny])
                    removed.append([nx, ny])
    return removed

def remove(removed):
    for i, j in removed:
        fields[i][j]='.'

def down():
    for i in range(6):
        for j in range(10, -1, -1):
            for k in range(11, j, -1):
                if fields[j][i] != '.' and fields[k][i] == '.':
                    fields[k][i] = fields[j][i]
                    fields[j][i] = "."
                    break

for _ in range(12):
    fields.append(list(map(str, input().strip('\n'))))

cnt = 0
while True:
    break_flag = True
    visited = [[False for _ in range(6)] for _ in range(12)]
    for i in range(12):
        for j in range(6):
            if fields[i][j] != '.':
                visited[i][j] = True
                remove_target=find(i,j)
                if len(remove_target) >= 4:
                    break_flag = False #내려갈거니까 한 바퀴 더 돌아서 검사필요
                    remove(remove_target)
    if break_flag == True:
        break
    down()
    cnt += 1

print(cnt)