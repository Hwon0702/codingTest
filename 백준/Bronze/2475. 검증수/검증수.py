import sys
input = sys.stdin.readline
sys.setrecursionlimit(10**6)
sum=0
numbers = list(map(int, input().split()))
for n in numbers:
    sum+=n **2
print(sum%10)