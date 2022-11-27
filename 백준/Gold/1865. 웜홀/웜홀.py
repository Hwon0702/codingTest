import sys
input = sys.stdin.readline
sys.setrecursionlimit(10**6)
def solution(n, graph):
    dist = [99999999 for _ in range(n+1)]
    dist[1] = 0
    for i in range(n):
        for j in range(1, n+1):
            for next, cost in graph[j]:
                if dist[next] > dist[j] + cost:
                    dist[next] = dist[j] + cost
                    if i == n-1:
                        return False
    return True


tc = int(input())
for test_case in range(tc):
    n, m, w = map(int, input().split())
    graph = [[] for _ in range(n+1)]
    for _ in range(m):
        s, e, c = map(int, input().split())
        graph[s].append([e,c])
        graph[e].append([s,c])

    for _ in range(w):
        s, e, c = map(int, input().split())
        graph[s].append([e,-c])
    if solution(n, graph):
        print("NO")
    else:
        print("YES")