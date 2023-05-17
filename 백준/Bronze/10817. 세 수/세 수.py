import sys 
input = sys.stdin.readline
sys.setrecursionlimit(10**6)
a = list(map(int, input().split()))
a.sort()
print(a[1])