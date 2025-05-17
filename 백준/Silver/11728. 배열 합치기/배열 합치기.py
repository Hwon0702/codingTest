import sys 
input = sys.stdin.readline
sys.setrecursionlimit(10**6)
n, m = map(int, input().split())
a = list(map(int, input().split()))
b = list(map(int, input().split()))
a = a+b
a.sort()
for v in a:
    print(v, end=" ")