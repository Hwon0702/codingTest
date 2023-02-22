import sys 
input = sys.stdin.readline 
sys.setrecursionlimit(10**6)

def factorial(n):
    num = 1
    for i in range(1, n+1):
        num *= i
    return num


tc = int(input())

for _ in range(tc):
    n, m = map(int, input().split())
    cnt = factorial(m) // (factorial(n) * factorial(m - n))
    print(cnt)