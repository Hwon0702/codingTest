import sys 
input = sys.stdin.readline
sys.setrecursionlimit(10**6)
n, m = map(int, input().split())
print(n//m)
print(n%m)