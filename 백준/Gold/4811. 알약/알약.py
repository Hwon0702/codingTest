import sys
input = sys.stdin.readline
sys.setrecursionlimit(10**6)

dp = [[0 for _ in range(31)] for i in range(31)] #최대 30일, 미리 만들어두기
for i in range(31):
    dp[0][i] = 1
for i in range(1,31):
    for j in range(i,31):
        dp[i][j] = dp[i-1][j] + dp[i][j-1]
while True:
    n = int(input())
    if n==0:
        break
    print(dp[n][n])
