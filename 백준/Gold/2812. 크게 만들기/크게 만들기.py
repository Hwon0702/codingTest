import sys 
input = sys.stdin.readline
sys.setrecursionlimit(10**6)
n,k = map(int, input().split())
nums = list(map(str, input().strip('\n')))
stack = []
for i in range(n):
    while k>0 and stack and stack[len(stack)-1] < nums[i]:
        stack.pop()
        k-=1
    stack.append(nums[i])

print(''.join(stack[:len(stack)-k])) 