import sys 
input = sys.stdin.readline 
sys.setrecursionlimit(10**6)
n = int(input())
m = int(input())
nums = list(map(int, input().split()))
nums.sort()
left = 0
right = n -1
cnt = 0
while left<right:
    if nums[left]+nums[right]==m:
        cnt+=1
        left+=1
        right-=1
    elif nums[left]+nums[right]>m:
        right-=1
    else:
        left+=1

print(cnt)