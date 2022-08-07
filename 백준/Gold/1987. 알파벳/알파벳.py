import sys
input = sys.stdin.readline
dx = [-1, 1, 0, 0]
dy = [0, 0, -1, 1]
def bfs():
    res=1
    queue=set([(0,0,graph[0][0])])
    while queue:
        x,y,visited=queue.pop()
        for i in range(4):
            nx=x+dx[i]
            ny=y+dy[i]
            if 0<=ny<h and 0<=nx<w and not graph[ny][nx] in visited:
                queue.add((nx,ny,visited+graph[ny][nx]))
                res=max(res,len(visited+graph[ny][nx]))
    return res
  
h, w=map(int, input().split())
graph=[]
for i in range(h):
    graph.append(list(input().strip('\n')))
print(bfs())