import sys 
input = sys.stdin.readline
sys.setrecursionlimit(10**6)
n = int(input())
res=[[0 for _ in range(10)] for _ in range(n+1)]
for i in range(1, 10):
    res[1][i]=1
for i in range(2, n+1):
    res[i][0]=res[i-1][1]
    res[i][9]=res[i-1][8]
    for j in range(1, 9):
        res[i][j]=res[i-1][j-1]+res[i-1][j+1]
sum = 0
for v in res[n]:
    sum+=v
print(sum%1000000000)
