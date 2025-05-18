import sys
input = sys.stdin.readline
sys.setrecursionlimit(10**6)

n = int(input())
res = 0
for _ in range(n):
    stack = []
    sArr = list(input().rstrip('\n'))
    for i in sArr:
        if not len(stack):
            stack.append(i)
        elif stack[-1] == i:
            stack.pop(-1)
        else:
            stack.append(i)

    if not len(stack):
        res += 1
print(res)