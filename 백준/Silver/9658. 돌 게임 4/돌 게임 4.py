import sys 
input = sys.stdin.readline
sys.setrecursionlimit(10**6)
dp = [False for _ in range(1001)]
n = int(input())
dp[2]=True
#1, 3, 4
for i in range(4, n+1):
    if not dp[i-1]:
        dp[i] = True
    if not dp[i-3]:
        dp[i] = True
    if not dp[i-4]:
        dp[i] = True
if not dp[n]:
    print("CY")
else:
    print("SK")
