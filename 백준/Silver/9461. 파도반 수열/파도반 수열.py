import sys
input = sys.stdin.readline
sys.setrecursionlimit(10**6)

results = [1]*(101)
tc = int(input())
for i in range(tc):
    N = int(input())
    if N<=3:
        print(1)
    else:
        if results[N]!=1:
            print(results[N])
        else:
            for i in range(4, N+1):
                results[i]=results[i-3]+results[i-2]
            print(results[N])