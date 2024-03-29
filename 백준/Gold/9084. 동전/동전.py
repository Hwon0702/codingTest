import sys
input = sys.stdin.readline
sys.setrecursionlimit(10**6)
tc = int(input())
for _ in range(tc):
    n = int(input())
    coins = list(map(int, input().split()))
    m = int(input())
    dp = [0 for _ in range(m+1)]
    dp[0] = 1
    for coin in coins:
        for i in range(m + 1):
            if i >= coin:
                dp[i] += dp[i - coin]
    print(dp[m])