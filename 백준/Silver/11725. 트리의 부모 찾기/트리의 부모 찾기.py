import sys
input = sys.stdin.readline
sys.setrecursionlimit(10**6)
n = int(input())
graph = [[] for i in range(n+1)]

for i in range(n-1):
    a, b = map(int, input().split())
    graph[a].append(b)
    graph[b].append(a)

tree = [0 for _ in range(n+1)]
def dfs(start):
    for i in graph[start]:
        if tree[i] == 0:
            tree[i] = start
            dfs(i)
dfs(1)
for x in range(2, n+1):
    print(tree[x])