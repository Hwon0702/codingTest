import sys 
input = sys.stdin.readline
sys.setrecursionlimit(10**6)
def find(n):
    if n==parents[n]:
        return n
    else:
        r= find(parents[n])
        dist[n]+=dist[parents[n]]
        parents[n] = r
    return parents[n]

# 합쳐서 최단거리로 
def union(x, y, k):
    root_x = parents[x]
    root_y = parents[y]
    if root_x != root_y:
        parents[root_y] = root_x
        dist[root_y] = (dist[x] + k) - dist[y]

while True:
    n, m =map(int,input().split())
    if n==0 and m==0:
        break

    parents = [i for i in range(n + 1)]
    dist = [0 for _ in range(n + 1)]
    for _ in range(m):
        tmp = list(input().split())
        a = int(tmp[1])
        b = int(tmp[2])
        find(a)
        find(b)
        if len(tmp)==4:
            r = int(tmp[3])
            union(a,b,r)
        elif len(tmp)==3:
            if parents[a] == parents[b]:
                print(dist[b] - dist[a])
            else:
                print("UNKNOWN")

