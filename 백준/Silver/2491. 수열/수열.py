import sys 
input = sys.stdin.readline
sys.setrecursionlimit(10**6)
n = int(input())
nums = list(map(int, input().split()))
dp = [[1 for _ in range(n)] for _ in range(2)]

for i in range(1,n):
    if nums[i-1]<=nums[i]: #증가
        dp[0][i] = dp[0][i-1]+1                         
    if nums[i-1]>=nums[i]: #감소
        dp[1][i] = dp[1][i-1]+1
print(max(map(max,dp)))