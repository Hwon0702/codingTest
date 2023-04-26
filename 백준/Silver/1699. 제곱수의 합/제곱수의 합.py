import sys 
input = sys.stdin.readline 
sys.setrecursionlimit(10**6)
squres = []

for i in range(1, int(100000**0.5)+1):
    squres.append(i**2)
def find(n):
    l, r = 0, len(squres)-1
    res = False
    while l<=r:
        if squres[l]+squres[r]==n:
            res= True
            break
        elif squres[l]+squres[r]>n:
            r-=1
        else:
            l+=1
    return res
def squre(a):
    if int(a**0.5)==a**0.5:
        return 1
    if find(a):
        return 2
    for v in squres: 
        if find(n-v):
            return 3
    else:
        return 4
    
n = int(input())
print(squre(n))