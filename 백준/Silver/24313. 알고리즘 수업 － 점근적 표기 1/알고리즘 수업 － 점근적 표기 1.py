import sys 
input = sys.stdin.readline
sys.setrecursionlimit(10**6)
a1, a2 = map(int, input().split())
c = int(input())
n = int(input())
if (a1*n + a2) <= (c*n) and a1 <= c:
    print(1)
else:
    print(0)
