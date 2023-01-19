import sys 
input = sys.stdin.readline
sys.setrecursionlimit(10**6)
n, m, c = map(int, input().split())
cost = [0]+list(map(int, input().split()))

parent = [i for i in range(n+1)] 
def find(x):
    if parent[x]==x:
        return x
    else:
        p = find(parent[x])
        parent[x]=p
        return p
def union(x, y):
    x, y = find(x), find(y)

    if x ==y :
        return
    elif cost[x]<=cost[y]:
        parent[y]=x 
    else:
        parent[x]=y

for _ in range(m):
    a, b = map(int, input().split())
    union(a, b)
res = 0
for i, r in enumerate(parent):
    if i==r:
        res+=cost[i]
if res<=c:
    print(res)
else:
    print("Oh no")

