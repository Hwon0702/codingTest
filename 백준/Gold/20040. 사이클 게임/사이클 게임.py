import sys 
input = sys.stdin.readline
sys.setrecursionlimit(10**6)
res= 0
def find(x):
    if parent[x] != x:
        parent[x] = find(parent[x])
    return parent[x]

def union(a, b, idx):
    global res
    a = find(a)
    b = find(b)
    if a != b:
        parent[max(a,b)] = min(a,b)
    elif res == 0:
        res = idx

n, m = map(int,input().split())
parent = list(range(n))
point_x = []
point_y = []
for idx in range(m):
    a, b = map(int, input().split())
    union(a, b, idx+1)
print(res)