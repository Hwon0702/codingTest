import sys 
from itertools import combinations
input = sys.stdin.readline
sys.setrecursionlimit(10**6)

strs = list(input().strip('\n'))
answer = set()
open = []
rem_target = []

for idx, word in enumerate(strs):
    if word == "(":
        open.append(idx)
    elif word == ")":
        rem_target.append((open.pop(), idx))

for i in range(1, len(rem_target) + 1):
    combs = combinations(rem_target, i)
    for single in combs:
        target = list(strs)
        for k in single:
            target[k[0]] = ""
            target[k[1]] = "" #pop으로 하면 인덱스 망가짐

        answer.add(''.join(target))

for ans in sorted(list(answer)):
    print(ans)