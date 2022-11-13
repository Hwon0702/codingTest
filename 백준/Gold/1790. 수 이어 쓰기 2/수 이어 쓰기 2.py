import sys 
input = sys.stdin.readline
sys.setrecursionlimit(10**6)
n, k = map(int, input().split())
mod = 1
cnt = 9
result = 0
while k > mod * cnt:
    k -= mod * cnt
    result += cnt
    mod += 1
    cnt *= 10
result += ((k - 1) // mod) + 1
if result > n:
    print(-1)
else:
    print(str(result)[(k - 1) % mod])