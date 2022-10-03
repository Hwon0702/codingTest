import sys 
input = sys.stdin.readline
sys.setrecursionlimit
target=input().strip('\n')
d = input().strip('\n')
cnt=0
idx = 0
while idx <= len(target) - len(d):
    if target[idx:idx + len(d)] == d:
        cnt += 1
        idx += len(d)
    else:
        idx += 1
print(cnt)