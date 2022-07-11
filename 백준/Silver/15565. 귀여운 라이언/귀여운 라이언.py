import sys
input = sys.stdin.readline
sys.setrecursionlimit(10**6)
N,K = map(int, input().split())
dolls = list(map(int, input().split()))
ryan = []
for i in range(0, N):
    if dolls[i]==1:
        ryan.append(i)
start, end = 0, K-1
min_length = 1000001
if len(ryan)<K:
    print(-1)
else:
    while start != end:
        ryan_count = ryan[end]-ryan[start]+1
        min_length = min(min_length, ryan_count)
        if end ==len(ryan)-1:
            break
        start+=1
        end+=1
    print(min_length)