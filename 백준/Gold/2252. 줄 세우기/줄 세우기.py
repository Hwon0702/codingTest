import sys 
from collections import deque
input = sys.stdin.readline
sys.setrecursionlimit(10**6)
n, m=map(int, input().split())
person = [[]for _ in range(n+1)]
cnt = [0]+[0 for _ in range(n)]
res=[]
for _ in range(m):
    be, af = map(int, input().split())
    person[be].append(af)
    cnt[af]+=1
q = deque()
for i in range(1, n+1):
    if cnt[i] == 0: 
        q.append(i)
        res.append(i)
while q:
    v = q.popleft()
    for i in person[v]:
        cnt[i] -=1
        if cnt[i] == 0:
            q.append(i)
            res.append(i)
for i in res:
    print(i, end=' ')