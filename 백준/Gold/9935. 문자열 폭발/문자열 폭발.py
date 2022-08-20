import sys
input=sys.stdin.readline
sys.setrecursionlimit(10**6)
str = list(input().strip('\n'))
r = input().strip('\n')
rem = list(r)
stack = []
for s in str:
    stack.append(s)
    if s ==rem[-1] and "".join(stack[-len(rem):])==r:
        del stack[-len(r):]
if len(stack)>0:
    print(''.join(stack))
else:
    print("FRULA")