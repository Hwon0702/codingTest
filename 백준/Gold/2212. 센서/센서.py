import sys 
input = sys.stdin.readline
sys.setrecursionlimit(10**6)
n = int(input())
k = int(input())
dist = []
nums = list(map(int, input().split()))
nums = sorted(nums)
for i in range(1, n):
    dist.append(nums[i]-nums[i-1])
dist = sorted(dist)
print(sum(dist[:n-k]))