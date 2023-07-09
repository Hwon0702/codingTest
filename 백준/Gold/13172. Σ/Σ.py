import sys 
input = sys.stdin.readline
sys.setrecursionlimit(10**6)

def power(num, t):
    if t == 1:
        return num % div
    elif (t % 2)==1:
        return num * power(num, t - 1) % div
    else:
        p = power(num, t // 2)
        return p * p % div

m = int(input())
div = 1000000007
res = 0

for _ in range(m):
    n, s = map(int, input().split())
    res += s * power(n, div - 2) % div
print(res % div)