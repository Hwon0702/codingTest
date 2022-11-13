import sys 
input = sys.stdin.readline
sys.setrecursionlimit(10**6)
n = int(input())
dp = [0 for _ in range(n + 1)]
squrt = [i * i for i in range(1, 317)] #N의 맥시멈 100,000의 제곱근 = 316.n
for i in range(1, n + 1):
    s = []
    for j in squrt:
        if j > i:
            break
        s.append(dp[i - j])
    dp[i] = min(s) + 1
print(dp[n])