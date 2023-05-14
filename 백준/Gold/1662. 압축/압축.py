import sys 
input = sys.stdin.readline 
sys.setrecursionlimit(10**6)

strs = input().rstrip('\n')
stack = []
cnt = 0
tmp = ''
for s in strs:
    if s == '(':
        stack.append((tmp, cnt-1))
        cnt = 0
    elif s == ')':
        zip, before = stack.pop()
        cnt = (int(zip)*cnt)+before
    else:
        cnt += 1
        tmp = s

print(cnt)