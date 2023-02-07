import sys 
input = sys.stdin.readline 
sys.setrecursionlimit(10**6)
n, m = map(int, input().split())
nums = list(map(int, input().split()))

left, right = 0, 1
cnt = 0
while right<=n and left<=right:
    s = sum(nums[left:right])
    if s == m:
        cnt += 1
        right += 1
    elif s < m:
        right += 1
    else:
        left += 1

print(cnt)
 