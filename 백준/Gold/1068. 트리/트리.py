from collections import deque
import sys 
input = sys.stdin.readline
sys.setrecursionlimit(10**6)

n = int(input())
tree = [[] for _ in range(n)]
origin = list(map(int, input().split()))
root = -1
for i in range(n):
    if origin[i] == -1:
        root = i
    else:
        tree[origin[i]].append(i)
deleted = int(input())
q = deque([root] if root != deleted else [])
count = 0
while q:
    v = q.popleft()
    leaf = 0
    for i in tree[v]:
        if i != deleted:
            leaf += 1
            q.append(i)
    if not leaf:
        count += 1
print(count)