import sys 
input = sys.stdin.readline 
sys.setrecursionlimit(10**6)
fibonacci = [0 for _ in range(46)]#1000000000 이 n의 맥시멈
fibonacci[1]=1
for i in range(2, 46):
    fibonacci[i] = fibonacci[i-1]+fibonacci[i-2]

tc = int(input())
for _ in range(tc):
    n = int(input())
    res = []
    while (n):
        for k in range(46):
            if (fibonacci[k] <= n):
                t = fibonacci[k]
        n = n - t
        res.append(t)
        res.sort()
    for v in res:
        print(v, end=' ')
    print()