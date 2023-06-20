import sys 
input = sys.stdin.readline 
sys.setrecursionlimit(10**6)
n = int(input())
a = map(int,input().split())
b, c = map(int, input().split())

res = 0
for i in a:
    i -= b
    cnt = 1
    if i > 0:
        cnt += int(i / c)
        if i % c != 0:
            cnt += 1
    res += cnt
print(res)