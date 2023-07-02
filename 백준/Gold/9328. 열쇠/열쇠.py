import sys
from collections import deque 
input = sys.stdin.readline 
sys.setrecursionlimit(10**6)

dx = [-1, 1, 0, 0]
dy = [0, 0, -1, 1]
def unlock(): 
    for i in range(h+2):
        for j in range(w+2):
            if graph[i][j].lower() in key_list:
                graph[i][j] = '.'
    key_list.clear()

def find():
    q = deque()
    visited = [[False for _ in range(w+2)]for _ in range(h+2)]
    q.append([0, 0])
    visited[0][0]=True
    global cnt
    if graph[0][0]=='$':
        cnt+=1
    while q:
        cx, cy = q.popleft()
        for i in range(4):
            nx = cx+dx[i]
            ny = cy+dy[i]
            if 0<=nx<h+2 and 0<=ny<w+2 and visited[nx][ny]==False and graph[nx][ny]!='*':
                if graph[nx][ny]=='$':
                    cnt+=1
                if ord('A') <= ord(graph[nx][ny]) <= ord('Z'):
                    continue
                if ord('a') <= ord(graph[nx][ny]) <= ord('z'):
                    key_list.append(graph[nx][ny])
                graph[nx][ny]='.'
                visited[nx][ny]=True
                q.append([nx, ny])
    return
tc = int(input())

for _ in range(tc):
    cnt = 0
    h,w=map(int,input().split()) 
    graph=[['.' for _ in range(w+2)]]
    for _ in range(h):
        row =  '.' +input().rstrip('\n')+ '.'
        graph.append(list(row))
    graph.append(['.' for _ in range(w+2)])
    key_list = deque(list(input().rstrip('\n')))    

    cnt = 0
    while key_list:
        if key_list[0] == '0': 
            key_list.clear() 
        if key_list: 
            unlock()
        find()  

    print(cnt)
