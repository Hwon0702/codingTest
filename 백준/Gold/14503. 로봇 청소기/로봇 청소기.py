import sys
input=sys.stdin.readline
sys.setrecursionlimit(10**6)
n, m = map(int, input().split())
r, c, d = map(int, input().split())
 
dx = [-1, 0, 1, 0]
dy = [0, 1, 0, -1]
 
graph = []
for i in range(n):
    graph.append(list(map(int, input().split())))
 
def turn_left(): #왼쪽으로 회전
    global d
    d = (d-1) % 4
 
count = 1 #처음은 벽이 아니므로, 1부터 시작
x, y = r, c
graph[x][y] = 2 # 방문처리
 
while True:
    check = False # 방문한 칸이 있는지 확인용
    for i in range(4): # 4방향
        turn_left()
        nx = x + dx[d]
        ny = y + dy[d]
        if 0 <= nx < n and 0 <= ny < m: #
            if graph[nx][ny] == 0: # 청소할 수 있는 칸
                count += 1
                graph[nx][ny] = 2 # 방문처리
                x, y = nx, ny
                check = True 
                break
    if not check: # 뒤로 가야함
        nx = x - dx[d]
        ny = y - dy[d]
        if 0 <= nx < n and 0 <= ny < m: 
            if graph[nx][ny] == 2: # 2라면 즉 이미 청소한 칸인경우 후진
                x, y = nx, ny
            elif graph[nx][ny] == 1: # 1인경우 즉 벽인 경우 
                print(count)
                break
        else:
            print(count)
            break
 
