import sys 
input = sys.stdin.readline 
sys.setrecursionlimit(10**6)
n = int(input())
res=0
nums = []
for i in range(1, n+1):
    if n%i==0:
        nums.append(i)
r = list(set(nums))
res = sum(r)
print(res*5-24)