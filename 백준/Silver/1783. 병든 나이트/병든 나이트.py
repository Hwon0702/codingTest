import sys 
input = sys.stdin.readline
sys.setrecursionlimit(10**6)
n, m = map(int, input().split())

if n == 1:
    print(1)
elif n == 2:
    print(min(4, (m + 1) // 2))
else:
    if m <= 6:
        print(min(4, m))
    else:
        print(m-2)