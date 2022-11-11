import sys
input = sys.stdin.readline

n, m = map(int, input().split())
graph = [[0 for _ in range(n)]for _ in range(n)]
for _ in range(m):
    s, e = map(int, input().split())
    graph[s-1][e-1]=1
#앞, 뒤 모두 있는지 확인(연결여부)
for path in range(n):
    for i in range(n):
        for j in range(n):
            if graph[i][path] == 1 and graph[path][j] == 1:
                graph[i][j] =1
res = 0
for i in range(n):
    check = 0
    for j in range(n):
        check += (graph[i][j] + graph[j][i]) #내 앞, 뒤가 모두 있을 경우
    # 자기 제외 모두 체크
    if check == n - 1:
        res +=1
print(res)
