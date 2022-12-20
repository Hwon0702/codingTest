import sys
input = sys.stdin.readline
sys.setrecursionlimit(10**6)
x, y = map(int, input().split()) 
dist = y - x 
if x==y:
    print(0)
else:
    n = 0
    while True:
        if dist <= n*(n+1):
            break
        n += 1
    if dist <= n**2:
        print(n*2-1)
    else:
        print(n*2)