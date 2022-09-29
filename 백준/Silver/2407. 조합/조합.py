import sys 
input = sys.stdin.readline
sys.setrecursionlimit(10**6)
n,m = map(int, input().split())

def factorial(n):
    res = 1
    for i in range(1, n+1):
        res = i*res 
    return res 

print(factorial(n)//(factorial(n-m)*factorial(m)))