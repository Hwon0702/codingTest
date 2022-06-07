import sys 
input = sys.stdin.readline
sys.setrecursionlimit(10**6)
n = int(input())
wines = [0]*(n+1)
res = [0]*(n+1)
for i in range(1, n+1):
    wines[i]=int(input())
if n==1:
    print(wines[1])
elif n==2:
    print(wines[1]+wines[2])
else:
    res[1]=wines[1]
    res[2]=wines[1]+wines[2]
    for i in range(3, n+1):
        res[i]= max(res[i-3]+wines[i-1]+wines[i], res[i-2]+wines[i], res[i-1])
    print(res[n])