import sys 
from collections import deque
input = sys.stdin.readline 
sys.setrecursionlimit(10**6)
n, k = map(int, input().split())
n_str = list(str(n))
max_n = -1
q = deque()
q.append([n, 0])
visited = set()
visited.add(chr(n)+','+chr(0))
while q:
    s, cnt = q.popleft()
    if cnt==k:
        max_n = max(max_n, s)
        continue
    n_str = list(str(s))
    for i in range(len(n_str)-1):
        for j in range(i+1, len(n_str)):
            if i==0 and n_str[j]=='0':
                continue
            n_str[i], n_str[j]=n_str[j], n_str[i]
            num = int(''.join(n_str))
            if chr(num)+','+chr(cnt+1) not in visited:
                q.append([num, cnt+1])
                visited.add(chr(num)+','+chr(cnt+1))
            n_str[j], n_str[i]=n_str[i], n_str[j]

print(max_n)