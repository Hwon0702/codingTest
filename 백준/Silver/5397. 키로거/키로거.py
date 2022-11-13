import sys
input = sys.stdin.readline
sys.setrecursionlimit(10**6)
tc = int(input())
res = []
for i in range(tc):
    s = list(input().strip('\n'))
    left = []
    right = []
    for i in s:
        if i == '<':
            if left:
                right.append(left.pop())
        elif i == '>':
            if right:
                left.append(right.pop())
        elif i == '-':
            if left:
                left.pop()
        else:
            left.append(i)
    left.extend(reversed(right))
    res.append(''.join(left))
for v in res:
    print(v)