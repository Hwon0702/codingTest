import math 
import sys 
input = sys.stdin.readline
sys.setrecursionlimit(10**6)
min_six=float('inf')
min_one=float('inf')
n, m = map(int, input().split())
for _ in range(m):
    s, o = map(int, input().split())
    min_six=min(s, min_six)
    min_one=min(o, min_one)
res=min(math.ceil(n/6)*min_six, min_six*(n//6)+(n-6*(n//6))*min_one, min_one*n)
print(res)
