import sys
input = sys.stdin.readline
sys.setrecursionlimit(10**6)
n, k = map(int,input().split())
a = list(map(int, input().split()))

res = []
res.append(sum(a[:k]))

for i in range(n - k):
    res.append(res[i] - a[i] + a[k+i])
        
print(max(res))