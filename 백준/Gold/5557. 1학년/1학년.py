import sys 
input = sys.stdin.readline
sys.setrecursionlimit(10**6)
n = int(input())
nums = list(map(int, input().split()))
dp = [[0 for _ in range(21)] for _ in range(n)]
dp[0][nums[0]] = 1
for i in range(1, n - 1):
    for j in range(21):
        if 0 <= j + nums[i] < 21: 
            dp[i][j + nums[i]] += dp[i - 1][j]
        if 0 <= j - nums[i] < 21: 
            dp[i][j - nums[i]] += dp[i - 1][j]
print(dp[n - 2][nums[-1]])