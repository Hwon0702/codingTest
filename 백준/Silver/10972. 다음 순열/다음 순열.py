import sys 
input = sys.stdin.readline 
sys.setrecursionlimit(10**6)
n = int(input())

nums = list(map(int, input().split()))
ans = -1
res = []
for i in range(n-2,-1,-1):
    if nums[i+1] > nums[i]:
        sub = 0
        for j in range(i+1,n):
            if nums[i] < nums[j]:
                sub = j
        nums[i], nums[sub] = nums[sub], nums[i]
        res += nums[:i+1]
        res += nums[i+1:][::-1]
        print(" ".join(map(str, res)))
        break

else:
    print(-1)