import sys
input = sys.stdin.readline
sys.setrecursionlimit(10**6)

N = int(input())
result = [0]*(1001)
result[0]=0
result[1]=1
result[2]=2
if N>=3:
    for i in range(3,N+1):
        result[i]=result[i-2]+result[i-1]
print(result[N]%10007)