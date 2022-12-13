import sys 
import math
input = sys.stdin.readline 
sys.setrecursionlimit(10**6)
n, m  = map(int, input().split())

frimes = [True for _ in range(1000001)]
for i in range(2, int(math.sqrt(m))+1):
    x = n//(i*i)
    if n%(i*i)!=0:
        x+=1
    while x*(i*i)<=m and 0<=x*(i*i)-n<m:
        frimes[x*(i*i)-n]=False
        x+=1
res = 0
for i in range(m-n+1):
    if frimes[i]:
        res+=1
print(res)
