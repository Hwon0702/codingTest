import sys 
from collections import deque
input = sys.stdin.readline
sys.setrecursionlimit(10**6)
n = int(input())
graph = [list(map(int, input().split())) for _ in range(n)]
res = [[0] * n for _ in range(n)]
res[0][0]=1
for x in range(n):
    for y in range(n):
        if x==n-1 and y==n-1:
            print(res[x][y])
            break
        nx = x+graph[x][y]
        ny = y+graph[x][y]
        if 0<=nx<n:
            res[nx][y]+=res[x][y]
        if 0<=ny<n:
            res[x][ny]+=res[x][y]