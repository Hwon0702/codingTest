import sys 
input = sys.stdin.readline 
dp = [1  for _ in range(1000001)]
s = [0 for _ in range(1000001)]

for i in range(2, 1000001):
    j=1
    while i*j < 1000001:
        dp[i*j] += i
        j += 1

for i in range(1, 1000001):
    s[i] = s[i-1] + dp[i]

n = int(input())
for _ in range(n):
    num=int(input())
    print(s[num])
