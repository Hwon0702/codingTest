import sys 
input = sys.stdin.readline 
n = int(input())
matrix = []
for _ in range(n):
    a, b = map(int, input().split())
    matrix.append([a, b])
dp = [[0 for _ in range(n)] for _ in range(n)]
for cnt in range(n-1):
    for i in range(n-1-cnt):
        j = i+cnt+1
        dp[i][j] = float('inf')
        for k in range(i, j):
            dp[i][j] = min(dp[i][j], dp[i][k] + dp[k+1][j] + matrix[i][0]*matrix[k][1]*matrix[j][1])
print(dp[0][-1])