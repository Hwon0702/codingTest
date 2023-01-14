import sys 
from collections import deque
input = sys.stdin.readline 
sys.setrecursionlimit(10**6)

def bfs():
    n = int(input())
    graph = [set() for _ in range(n+1)]
    graph[0].add(1)
    for _ in range(n-1):
        s, e = map(int, input().split())
        graph[s].add(e)
        graph[e].add(s)
    q = deque()
    q.append(0)
    idx = 0 
    ans = list(map(int, input().split()))
    for x in ans:
        while x not in graph[q[idx]]:
            idx+=1
            if idx==len(q):
                return 0
        q.append(x)
    return 1 
print(bfs())
