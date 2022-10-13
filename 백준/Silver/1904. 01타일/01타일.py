import sys 
input = sys.stdin.readline
sys.setrecursionlimit(10**6)
n = int(input())
res = [0 for _ in range(n+1)]
if n<=2:
    print(n)
else:
    res[1]=1
    res[2]=2
    for i in range(3, n+1):
        res[i]=(res[i-1]+res[i-2])%15746
    print(res[n])