import sys 
input= sys.stdin.readline
sys.setrecursionlimit(10**6)

t, w = map(int,input().split())
dp = [[0 for _ in range(w+1)] for _ in range(t)]
tree = [int(input())for _ in range(t)]
for i in range(t):
    if tree[i]==1:
        dp[i][0]=dp[i-1][0]+1
    else:
        dp[i][0]=dp[i-1][0]
    for j in range(1, w+1):
        if tree[i]==1 and j%2==0:
            dp[i][j]=max(dp[i-1][j], dp[i-1][j-1])+1
        elif tree[i]==2 and j%2==1:
            dp[i][j]=max(dp[i-1][j], dp[i-1][j-1])+1
        else:
            dp[i][j]=max(dp[i-1][j], dp[i-1][j-1])
print(max(dp[-1]))