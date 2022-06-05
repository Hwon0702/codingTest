import sys
input = sys.stdin.readline
sys.setrecursionlimit(10**6)
n = int(input())
if n==0:
    print(0)
elif n==1:
    print(1)
else:
    fib_arr = [0 for i in range(n+1)]
    fib_arr[1]=1
    for i in range(2, (n+1)):
        fib_arr[i] = fib_arr[i-2]% 1000000007+fib_arr[i-1]%1000000007
    print(fib_arr[n]%1000000007)