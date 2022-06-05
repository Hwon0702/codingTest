import sys
from unittest import result
input = sys.stdin.readline
sys.setrecursionlimit(10**6)
n=int(input())
result_arr = [0]*(n+1)

if n==1:
    print(1)
elif n==2:
    print(3)
else:
    result_arr[1]=1
    result_arr[2]=3
    for i in range(3, n+1):
        result_arr[i]=(result_arr[i-2]*2+result_arr[i-1])%10007
    print(result_arr[n]%10007)