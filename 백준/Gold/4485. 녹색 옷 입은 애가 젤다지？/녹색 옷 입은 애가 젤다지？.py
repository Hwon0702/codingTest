from cmath import cos
import heapq
import sys 

input = sys.stdin.readline
sys.setrecursionlimit(10**6)
dx = [1, -1, 0, 0]
dy = [0, 0, 1, -1]
def dijkstra(num, graph):
    cost = [[float('inf') for _ in range(len(graph))] for _ in range(len(graph))]
    q = []
    heapq.heappush(q, [graph[0][0], 0, 0])
    cost[0][0]=0
    while q:
        c, x, y = heapq.heappop(q)
        if x == n-1 and y ==n-1:
            print("Problem {0}: {1}".format(num, c))
            return
        for i in range(4):
            nx = x+dx[i]
            ny = y+dy[i]
            if 0<=nx<len(graph) and 0<=ny<len(graph):
                ncost = c+graph[nx][ny]
                if ncost<cost[nx][ny]:
                    cost[nx][ny] = ncost
                    heapq.heappush(q, [ncost, nx, ny])
num = 1
while True:
    n = int(input())
    if n==0:
        break
    graph = [list(map(int, input().split())) for _ in range(n)]
    dijkstra(num, graph)
    num+=1
