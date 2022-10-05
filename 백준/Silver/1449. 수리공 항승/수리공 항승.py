import sys 
input = sys.stdin.readline
sys.setrecursionlimit(10**6)
N, L = map(int, input().split())
leak = list(map(int, input().split()))
leak.sort()
start = leak[0]
end = leak[0] + L
cnt = 1
for i in range(N):
    if start <= leak[i] < end:
        continue
    else:
        start = leak[i]
        end = leak[i] + L
        cnt += 1
print(cnt)