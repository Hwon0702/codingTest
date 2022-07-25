import sys 
input = sys.stdin.readline
sys.setrecursionlimit(10**6)
N, K= map(int, input().split())
weights = [0]
values = [0]
for _ in range(N):
    w, v = map(int, input().split())
    weights.append(w)
    values.append(v)

dp=[[0  for i in range(K+1)] for k in range(N+1)]
for w in range(1, N+1):

    for i in range(1, K+1):
        if i>= weights[w]:
                dp[w][i]=max(dp[w-1][i], dp[w-1][i-weights[w]]+values[w])
        else:
            dp[w][i]=dp[w-1][i]
print(dp[N][K])