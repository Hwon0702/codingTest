import sys 
from copy import deepcopy
input = sys.stdin.readline
sys.setrecursionlimit(10**6)
n = int(input())
graph = [list(map(int, input().split())) for _ in range(n)]
res = 0


def up(graph):
    for j in range(n):
        target = 0

        for i in range(1, n):
            if graph[i][j]:
                tmp = graph[i][j]
                graph[i][j] = 0

                if graph[target][j] == 0:
                    graph[target][j] = tmp

                elif graph[target][j] == tmp:
                    graph[target][j] = graph[target][j]*2
                    target += 1

                else:
                    target += 1
                    graph[target][j] = tmp

    return graph


def down(graph):
    for j in range(n):
        target = n - 1

        for i in range(n - 2, -1, -1):
            if graph[i][j]:
                tmp = graph[i][j]
                graph[i][j] = 0

                if graph[target][j] == 0:
                    graph[target][j] = tmp
                elif graph[target][j] == tmp:
                    graph[target][j] =graph[target][j]*2
                    target -= 1
                else:
                    target -= 1
                    graph[target][j] = tmp

    return graph


def left(graph):
    for i in range(n):
        target = 0

        for j in range(1, n):
            if graph[i][j]:
                tmp = graph[i][j]
                graph[i][j] = 0

                if graph[i][target] == 0:
                    graph[i][target] = tmp
                elif graph[i][target] == tmp:
                    graph[i][target] = graph[i][target]*2
                    target += 1
                else:
                    target += 1
                    graph[i][target] = tmp

    return graph


def right(graph):
    for i in range(n):
        target = n - 1

        for j in range(n - 2, -1, -1):
            if graph[i][j]:
                tmp = graph[i][j]
                graph[i][j] = 0

                if graph[i][target] == 0:
                    graph[i][target] = tmp
                elif graph[i][target] == tmp:
                    graph[i][target] *= 2
                    target -= 1
                else:
                    target -= 1
                    graph[i][target] = tmp

    return graph


def dfs(graph, cnt):
    global res

    if cnt == 5:
        res = max(res, max(map(max, graph)))
        return

    dfs(left(deepcopy(graph)), cnt + 1)
    dfs(right(deepcopy(graph)), cnt + 1)
    dfs(up(deepcopy(graph)), cnt + 1)
    dfs(down(deepcopy(graph)), cnt + 1)


dfs(graph, 0)
print(res)