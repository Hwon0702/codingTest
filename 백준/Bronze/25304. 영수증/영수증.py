import sys
input = sys.stdin.readline
sys.setrecursionlimit(10**6)

x = int(input())
n = int(input())
sum = 0

for _ in range(n):
    p,c = map(int, input().split())
    sum+=p*c

if sum==x:
    print("Yes")
else:
    print("No")