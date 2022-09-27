import sys 
input = sys.stdin.readline
sys.setrecursionlimit(10**6)
n = int(input())
numbers = list(map(int, input().split()))
nums = [0]
start, end, mid = 0,0,0
for num in numbers:
    if nums[-1]<num:
        nums.append(num)
    else:
        start = 0
        end = len(nums)
        while start<end:
            mid = (start+end)//2
            if nums[mid]<num:
                start=mid+1
            else:
                end = mid
        nums[end]=num
print(len(nums)-1)