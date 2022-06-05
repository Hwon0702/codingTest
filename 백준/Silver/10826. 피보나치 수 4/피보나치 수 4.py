import sys
input = sys.stdin.readline
sys.setrecursionlimit(10**6)

a=1
b=1
n = int(input())
result = 0
if n==0:
    print(0)
elif n>0 and n<=2:
    print(1)
else:
    for i in range(n-2):
        result = a+b 
        a=b 
        b=result
    print(result)
