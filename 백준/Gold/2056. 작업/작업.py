import sys 
input = sys.stdin.readline
sys.setrecursionlimit(10**6)
n = int(input())
dp = [0 for _ in range(n+1)]
graph = [[] for _ in range(n+1)]
for i in range(1, n+1):
    jobs = list(map(int, input().split()))
    cost = jobs[0]
    m = jobs[1] 
    dp[i] = cost
    jobs = jobs[2:]
    for job in jobs:
        graph[i].append(job)

for i in range(1, n+1):
    tmp = 0
    for j in graph[i]:
        tmp = max(tmp, dp[j])
    dp[i] += tmp
print(max(dp))