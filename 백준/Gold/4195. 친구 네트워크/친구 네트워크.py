import sys
input = sys.stdin.readline
sys.setrecursionlimit(10**6)
def find(x):
    if parent[x] == x: 
        return x
    else:
        p = find(parent[x])
        parent[x] = p
        return p

def union(x, y):
    x, y = find(x), find(y)

    if x != y:
        parent[y] = x
        number[x] += number[y]
    print(number[x])


tc = int(input())
for _ in range(tc):
    num = int(input())
    parent, number = {}, {}
    for i in range(num):
        s, e = map(str, input().split())
        if s not in parent:
            parent[s] = s
            number[s] = 1
        if e not in parent:
            parent[e] = e
            number[e] = 1
        union(s, e)