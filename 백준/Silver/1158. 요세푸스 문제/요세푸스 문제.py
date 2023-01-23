import sys 
from collections import deque
input = sys.stdin.readline
sys.setrecursionlimit(10**6)
n, k = map(int, input().split())
people = deque()
for i in range(1, n+1):
    people.append(i)
res = []
cnt = 0
while people:
    c = people.popleft()
    cnt+=1
    if cnt == k:
        res.append(c)
        cnt = 0
        continue
    people.append(c)
print('<'+', '.join(map(str, res))+'>')
