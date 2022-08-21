import sys
from collections import deque
input=sys.stdin.readline
sys.setrecursionlimit(10**6)
n = int(input())
nums = list(map(int, input().split()))
good = 0
nums.sort()
for i in range(n):
    targets = nums[:i] + nums[i+1:]
    left, right = 0, len(targets)-1
    while left < right:
        sum = targets[left] + targets[right]
        if sum == nums[i]:
            good += 1
            break
        elif sum < nums[i]:
            left += 1
        else:
            right -= 1
print(good)