from math import factorial
import sys
input = sys.stdin.readline
sys.setrecursionlimit(10**6)
n, k = map(int, input().split())
res = factorial(n) // (factorial(k) * factorial(n - k))
print(res)