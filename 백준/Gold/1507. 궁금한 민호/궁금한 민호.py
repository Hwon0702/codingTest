import sys 
input = sys.stdin.readline
sys.setrecursionlimit(10**6)
def floyd_warshall():
    ignore = [[False for _ in range(n)] for _ in range(n)]
    for k in range(n):
        for a in range(n):
            for b in range(n):
                if a==b or a==k or b==k:
                    continue
                if graph[a][b]==graph[a][k]+graph[k][b]:#최단경로 이미 나옴
                    ignore[a][b]=True #무시함
                elif graph[a][b]>graph[a][k]+graph[k][b]: #길이 없음
                    print(-1)
                    exit()
    return ignore
n = int(input())
graph = []
for _ in range(n):
    graph.append(list(map(int, input().split())))
res = 0
ignore = floyd_warshall()
for i in range(n):
    for j in range(i+1, n): #같은 값이 반복해서 나옴 -> 절반만
        if not ignore[i][j]:
            res+=graph[i][j]
print(res)