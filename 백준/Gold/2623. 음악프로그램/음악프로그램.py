import sys
from collections import deque
input = sys.stdin.readline
sys.setrecursionlimit(10**6)
n, m = map(int, input().split())
singers = [[]for _ in range(n+1)]
cnt = [0 for _ in range(n+1)]
q = deque()
res = []
cnt[0]=-1
for _ in range(m):
    pd = list(map(int, input().split()))
    for i in range(1,len(pd)-1):
        singers[pd[i]].append(pd[i+1])
        cnt[pd[i+1]]+=1 #다음에 연결된 값에 +1
for i in range(1,n+1):
    if cnt[i] == 0: #연결된 값이 없다면 -> 순서 배열(res)에 넣고 연결된 다음 노드로 감. 연결 순서대로 큐에 넣기
        q.append(i)
        res.append(i)
while q:
    v = q.popleft()
    for i in singers[v]: #큐에 들어가 있는 순서대로 확인
        cnt[i] -=1 #카운트 감소 -> 하나씩 빼가면서 확인
        if cnt[i] == 0:
            q.append(i)  #연결된 값이 없다면 -> 순서 배열(res)에 넣고 연결된 다음 노드로 감. 연결 순서대로 큐에 넣기
            res.append(i)

if len(res)==n:
    for _, v in enumerate(res):
        print(v)
else:
    print(0)