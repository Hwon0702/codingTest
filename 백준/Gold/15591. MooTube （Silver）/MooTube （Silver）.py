from collections import deque
import sys 
input = sys.stdin.readline 

def find(k, v):
    global N, Q, graph
    visited = [False for _ in range(N)]
    queue = deque()
    queue.append([v, float('inf')])
    visited[v]=True
    res = 0
    while queue:
        now, cost = queue.popleft()
        for next, n_cost in graph[now].items():
            next_cost = min(n_cost, cost)
            if next_cost>k and visited[next]==False:
                visited[next]=True
                res+=1
                queue.append([next, next_cost])
    return res

N, Q = map(int, input().split())
graph = [{}for _ in range(N)]

for i in range(N-1):
    graph[i]={}
for _ in range(N-1):
    p, q, r = map(int, input().split())
    graph[p-1][q-1] = r
    graph[q-1][p-1]=r
for _ in range(Q):
    k, v = map(int, input().split())
    print(find(k-1, v-1))

