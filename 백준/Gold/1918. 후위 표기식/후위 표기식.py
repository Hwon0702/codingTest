import sys 
input = sys.stdin.readline
sys.setrecursionlimit(10**6)

res = ''
stack = []
calculate = list(map(str, input().strip('\n')))
methods = ['+', '-']
priority_methods = ['*', '/']
for v in calculate:
    if not v in methods and not v in priority_methods and v!='(' and v!=')':
        res+=v
    else:
        if v =='(':
            stack.append(v)
        elif v in priority_methods:
            while stack and (stack[-1] in priority_methods):
                res+=stack.pop()
            stack.append(v)
        elif v in methods:
            while stack and stack[-1] !='(':
                res+=stack.pop()
            stack.append(v)
        elif v==')':
            while stack and stack[-1]!='(':
                res+=stack.pop()
            stack.pop()
while stack:
    res+=stack.pop()
print(res)
