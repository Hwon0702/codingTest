import sys 
from collections import deque
input = sys.stdin.readline 
sys.setrecursionlimit(10**6)

n = int(input())
tree = [{} for _ in range(n+1)]
max_node = 0
def find(start):
    global max_node
    visited = [False for _ in range(n+1)]
    visited[0]=True
    visited[start]=True
    q = deque()
    q.append([start, 0])
    max_len = 0
    
    while q:
        now, c = q.popleft()
        for dest, cost in tree[now].items():
            if visited[dest]==False:
                visited[dest]=True
                q.append([dest,c+cost])
                if max_len < c+cost:
                    max_len = c+cost
                    max_node = dest
    return max_len
for i in range(n+1):
    tree[i]={}
for _ in range(n-1):
    p, s, c = map(int, input().split())
    tree[p][s]=c 
    tree[s][p]=c
find(1)
print(find(max_node))