from collections import deque
import sys
input = sys.stdin.readline
n = int(input())
tree  = [[]for _ in range(n+1)]
depth = [0] * (n+1)
parant = [0] * (n+1)
visited = [False] * (n+1)
for _ in range(n-1):
    s,e = map(int, input().split())
    tree[s].append(e)
    tree[e].append(s)

def bfs(s):
    queue = deque()
    queue.append(s)
    while queue:
        node = queue.popleft()
        visited[node] = True
        for i in tree[node]:
            if not visited[i]:
                depth[i] = depth[node] + 1
                parant[i] = node
                queue.append(i)

def lca(a,b):
    if depth[a] < depth[b]:
        temp = a
        a=b
        b=temp
    diff = depth[a] - depth[b]
    for _ in range(diff):
        a = parant[a]
    while a!=b:
        a = parant[a]
        b = parant[b]
    return a

bfs(1)

m = int(input())
for _ in range(m):
    a,b = map(int, input().split())
    print(lca(a, b))