import sys
input = sys.stdin.readline
sys.setrecursionlimit(10**6)
n = int(input())
results = [1]*(n+1)
if n<=2:
    print(1)
else:
    for i in range(3, n+1):
        results[i]=results[i-2]+results[i-1]
    print(results[n])