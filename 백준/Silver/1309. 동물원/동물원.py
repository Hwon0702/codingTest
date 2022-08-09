import sys 
input = sys.stdin.readline
sys.setrecursionlimit(10**6)
n = int(input())
res = [[0,0,0] for _ in range(n+1)]
res[1][0]=1
res[1][1]=1
res[1][2]=1
for i in range(2, n+1):
    res[i][0] = (res[i - 1][0] + res[i - 1][1] + res[i - 1][2]) % 9901
    res[i][1] = (res[i - 1][0] + res[i - 1][2]) % 9901
    res[i][2] = (res[i - 1][0] + res[i - 1][1]) % 9901
print((res[n][0]+res[n][1]+res[n][2])%9901)