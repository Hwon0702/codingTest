import sys 
input = sys.stdin.readline
sys.setrecursionlimit(10**6)

start = list(map(str, input().strip('\n')))
target = list(map(str, input().strip('\n')))
flag = 0

while target:
    if target[-1]=='A':
        target.pop()
    elif target[-1]=='B':
        target.pop()
        target.reverse()
    if target==start:
        flag=1
        break

print(flag)
