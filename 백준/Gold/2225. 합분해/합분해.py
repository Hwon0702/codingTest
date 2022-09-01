import sys
input = sys.stdin.readline
sys.setrecursionlimit(10**6)
n, k = map(int, input().split())
res = [[1]*(n+1) for _ in range(k+1)]
for i in range(2, k+1):
    for j in range(1, n+1):
        res[i][j]=(res[i-1][j]+res[i][j-1])% 1000000000
print(res[k][n])