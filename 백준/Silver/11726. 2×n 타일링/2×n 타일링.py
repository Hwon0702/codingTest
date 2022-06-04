import sys
input = sys.stdin.readline
sys.setrecursionlimit(10**6)

N = int(input())
result = [0]*(1001) #  1넣으면 에러나서 그냥 최대값 잡아버림
result[0]=0
result[1]=1
result[2]=2
if N>=3:
    for i in range(3,N+1):
        result[i]=result[i-2]%10007+result[i-1]%10007
print(result[N]%10007)