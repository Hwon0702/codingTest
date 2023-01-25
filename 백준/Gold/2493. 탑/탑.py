import sys 
input = sys.stdin.readline
sys.setrecursionlimit(10**6)
n = int(input())
towers = list(map(int, input().split()))
stack = []
res = [0]
stack.append([1, towers[0]])

for i in range(1, n):
    while stack:
        if stack[-1][1] >= towers[i]:
            res.append(stack[-1][0])
            stack.append([i + 1, towers[i]])
            break
        
        else:
            stack.pop()

    if not stack:
        res.append(0)
        stack.append([i + 1, towers[i]])

print(*res)