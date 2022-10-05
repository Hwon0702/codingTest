import sys 
input = sys.stdin.readline
sys.setrecursionlimit(10**6)
n = int(input())
f = int(input())
mod = (n//100)*100
n = (n//100)*100
while mod%f!=0:
    mod+=1
print(format(mod%100, '02'))