import sys 
from collections import deque
from itertools import combinations
input = sys.stdin.readline 
sys.setrecursionlimit(10**6)
n = int(input())
q = deque()
people = [0]+list(map(int, input().split()))
graph = [{} for _ in range(n+1)]
result = float('inf')
for i in range(1, n+1):
    graph[i] = set()
    tmp = list(map(int, input().split()))
    for j in tmp[1:]:
        graph[i].add(j)
def calculate(nodes):
    q = deque()
    visited = set()
    start = nodes[0]
    q.append(start)
    visited.add(start)
    total = 0
    while q:
        current = q.popleft()
        total+= people[current]
        for next in graph[current]:
            if next not in visited and next in nodes:
                q.append(next)
                visited.add(next)
    return total, len(visited)

for i in range(1, n//2 + 1):
    combination_list = list(combinations(range(1, n+1), i))
    for combi in combination_list:
        sum1, v1 = calculate(combi) #선거구역1
        sum2, v2 = calculate([i for i in range(1, n+1) if i not in combi]) #선거구역2
        if v1 + v2 == n: #선거구역1+2 = n일 때만 검사
            result = min(result, abs(sum1 - sum2))

if result != float('inf'):
    print(result)
else:
    print(-1)