import sys 
input = sys.stdin.readline
sys.setrecursionlimit(10**6)
min_num = float('inf')
max_num = -float('inf')
num = int(input())
nums= list(map(int, input().split()))
methods = list(map(int, input().split()))

def dfs(depth, res, plus, minus, multiple, divide):
    global min_num, max_num
    if depth==num:
        min_num=min(min_num, res)
        max_num=max(max_num, res)
        return
    
    if plus:
        dfs(depth+1, res+nums[depth], plus-1, minus, multiple, divide)
    if minus:
        dfs(depth+1, res-nums[depth], plus, minus-1, multiple, divide)
    if multiple:
        dfs(depth+1, res*nums[depth], plus, minus, multiple-1, divide)
    if divide:
        dfs(depth+1, int(res/nums[depth]), plus, minus, multiple, divide-1)

dfs(1, nums[0], methods[0], methods[1], methods[2], methods[3])

print(max_num)
print(min_num)