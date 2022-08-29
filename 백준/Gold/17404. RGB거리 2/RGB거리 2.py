import sys 
from collections import deque
input = sys.stdin.readline
sys.setrecursionlimit(10**6)

n = int(input())
cost = [list(map(int, input().split())) for _ in range(n)]
res = [[0 for _ in range(3)] for _ in range(n)]
ans = 9999999999
for f in range(3): 
    for i in range(3):
        res[0][i] = 9999999999 if i == f else cost[0][i] #같은게 선택됐을 때
    for i in range(1, n):
        res[i][0] = min(res[i-1][1],res[i-1][2]) + cost[i][0]
        res[i][1] = min(res[i-1][0],res[i-1][2]) + cost[i][1]
        res[i][2] = min(res[i-1][1],res[i-1][0]) + cost[i][2]

    for i in range(3):
        if i == f:
            ans = min(ans, res[n-1][i])

print(ans)