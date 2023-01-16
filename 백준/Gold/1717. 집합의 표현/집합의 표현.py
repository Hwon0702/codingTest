import sys 
input = sys.stdin.readline 
sys.setrecursionlimit(10**6)
n, m = map(int, input().split())

parent = [0] * (n+1) 
for i in range(n+1): 
    parent[i] = i

def find(x):
    if parent[x] == x: 
        return x
    else:
        p = find(parent[x])
        parent[x] = p
        return p

def union(x, y):
    x, y = find(x), find(y)

    if x == y:
        return
    elif x<y:
        parent[y]=x 
    else:
        parent[x]=y
for _ in range(m):
    op, a, b = map(int, input().split())
    if op==0:
        union(a, b)
    elif op==1:
        if find(a)==find(b):
            print('YES')
        else:
            print('NO')
