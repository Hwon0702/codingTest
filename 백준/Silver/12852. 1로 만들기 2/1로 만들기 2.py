import sys 
from collections import deque
input = sys.stdin.readline 
n = int(input())
res = []
q = deque()
q.append([n])
while q:
    history = q.popleft()
    num = history[-1]
    if num==1:
        res = history
        break
    if num%3==0:
        q.append(history+[int(num/3)])
    if num%2==0:
        q.append(history+[int(num/2)])
    q.append(history+[int(num-1)])
print(len(res)-1)
print(*res)
