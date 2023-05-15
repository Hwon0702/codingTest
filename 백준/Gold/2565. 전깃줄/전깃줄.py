import sys 
input = sys.stdin.readline 
sys.setrecursionlimit(10**6)
n = int(input())
lines = []
dp = [1 for _ in range(n)]
for _ in range(n):
    s, e = map(int, input().split())
    lines.append([s, e])
lines.sort()

for i in range(1, n):
  for j in range(0, i):
    if lines[j][1] < lines[i][1]:
      dp[i] = max(dp[i], dp[j]+1)
print(n-max(dp))