import sys 
input = sys.stdin.readline
sys.setrecursionlimit(10**6)
n = int(input())
nums = list(map(int, input().split()))
nums = sorted(nums)
compare = 1
for v in nums:
    if compare<v:
        break
    compare+=v
print(compare)
