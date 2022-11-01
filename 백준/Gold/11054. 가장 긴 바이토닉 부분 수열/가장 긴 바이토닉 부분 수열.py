import sys 
input = sys.stdin.readline
sys.setrecursionlimit(10**6)

n = int(input())
nums = list(map(int, input().split()))
reverse = nums[:]
reverse.reverse()
increase = [1 for _ in range(n)]
decrease = [1 for _ in range(n)]
for i in range (n):
    for j in range (i):
        if nums[i]>nums[j]:
            increase[i] = max(increase[i],increase[j]+1)
        if reverse[i]>reverse[j]:
            decrease[i] =max(decrease[i], decrease[j]+1)
res = [0 for _ in range(n)]
for i in range(n):
    res[i] = increase[i] + decrease[n-i-1] -1
print(max(res))