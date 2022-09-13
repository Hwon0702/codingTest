import sys 
input = sys.stdin.readline
sys.setrecursionlimit(10**6)

def floyd_warshal(bus, n):
    cost  = bus
    for k in range(1, n + 1):
        cost[k][k]=0
        for a in range(1, n + 1):
            for b in range(1, n + 1):
                cost[a][b] = min(bus[a][b], bus[a][k] + bus[k][b])
    return cost
n = int(input())
m = int(input())
bus = [[float('inf')] * (n+1) for _ in range(n+1)]
for _ in range(m):
    s, e, c = map(int, input().split())
    bus[s][e] = min(bus[s][e], c)

res = floyd_warshal(bus, n)
for a in range(1, n + 1):
    for b in range(1, n + 1):
        if res[a][b] == float('inf'):
            print(0, end=" ")
        else:
            print(res[a][b], end=" ")
    print()