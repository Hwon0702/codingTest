import sys 
input = sys.stdin.readline
sys.setrecursionlimit(10**6)
n, m = map(int, input().split())
nums = list(map(int, input().split()))
nums.append(0)
mod = [0 for _ in range(m)]

for i in range(n):
    nums[i] += nums[i - 1]
    mod[nums[i] % m] += 1
cnt = mod[0]
for i in mod:
    cnt += i * (i - 1) // 2
print(cnt)