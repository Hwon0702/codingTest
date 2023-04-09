import sys 
from collections import deque
input = sys.stdin.readline
sys.setrecursionlimit(10**6)

strs = list(input().rstrip())
length = len(strs)
checked = [[False for _ in range(length)] for _ in range(length)]
dp = [0 for _ in range(length)]
for i in range(length):
    for j in range(length-i):
        if i == 0:
            checked[j][j] = 1
            continue
        k = i+j
        if strs[j] == strs[k] and (checked[j+1][k-1] or j==k-1): #팰린드롬이면
            checked[j][k] = 1
for i in range(length): #팰린드롬 확인
	for j in range(i+1):
		if checked[j][i]:
			if dp[i] == 0 or dp[i]>dp[j-1]+1:
				dp[i] = dp[j-1]+1
print(dp[length-1])