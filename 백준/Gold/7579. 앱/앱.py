import sys 
input = sys.stdin.readline
sys.setrecursionlimit(10**6)
n, m = map(int, input().split())
active = [0]+list(map(int, input().split()))
inactive =  [0]+list(map(int, input().split()))
dp=[[0 for _ in range(sum(inactive)+1)] for _ in range(n+1)]
result = float('inf')
for i in range(1, n+1):
    for j in range(1, sum(inactive)+1):
        if j<inactive[i]:
            dp[i][j]=dp[i-1][j]
        else:
            dp[i][j]= max(dp[i - 1][j - inactive[i]] + active[i], dp[i - 1][j])
        if dp[i][j] >= m:
            result = min(result, j)
if m != 0:
    print(result)
else:
    print(0)
