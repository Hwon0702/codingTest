import sys 
input = sys.stdin.readline 
sys.setrecursionlimit(10**6)
n = int(input())
dp = [False for _ in range(1001)]
dp[2]=True

for i in range(6, n+1):
    dp[i] = not dp[i-1] and not dp[i-3] and not dp[i-4]
if dp[n]:
    print("CY")
else:
    print("SK")