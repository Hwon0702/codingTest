import sys
input = sys.stdin.readline 
sys.setrecursionlimit(10**6)
tc = int(input())

for _ in range(tc):
    n = int(input())
    parent = [0 for _ in range(n + 1)]

    for _ in range(n - 1):
        a, b = map(int, input().split())
        parent[b] = a

    a, b = map(int, input().split())
    targetA = [a]
    targetB = [b]
    while parent[a]:
        targetA.append(parent[a])
        a = parent[a]

    while parent[b]:
        targetB.append(parent[b])
        b = parent[b]

    levelA = len(targetA) - 1
    levelB = len(targetB) - 1

    while targetA[levelA] == targetB[levelB]:
        levelA -= 1
        levelB -= 1

    print(targetA[levelA + 1])