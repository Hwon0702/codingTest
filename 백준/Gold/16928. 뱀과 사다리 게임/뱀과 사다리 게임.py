import sys 
from collections import deque
input = sys.stdin.readline
sys.setrecursionlimit(10**6)
n, m = map(int, input().split())
ladder = []
snake = []
for _ in range(n):
    ladder.append(list(map(int, input().split())))

for _ in range(m):
    snake.append(list(map(int, input().split())))

q = deque()
q.append([1, 0])
visited=[False for _ in range(101)]
visited[1]=True
while q:
    now, cnt = q.popleft()
    if now==100:
        print(cnt)
        break
    for j in range(1, 7):
        next = now+j
       
        if next>100 or visited[next]:
            continue

        for l in ladder:
            if l[0]==next:
                next = l[1]
        for s in snake:
            if s[0]==next:
                next=s[1]
        
        visited[next]=True
        q.append([next,cnt+1])
