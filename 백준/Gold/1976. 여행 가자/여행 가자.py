import sys 
input = sys.stdin.readline
sys.setrecursionlimit(10**6)
n = int(input())
m = int(input())
graph = [] 
def floyd_warshal(graph, n):
    connect = graph
    for k in range(n):
        for a in range(n):
            for b in range(n):
                if a==b:
                    connect[a][b]=1
                if graph[a][k]==1 and graph[k][b]==1:
                    connect[a][b]=1
    return connect
for _ in range(n):
    graph.append(list(map( int,input().split())))
race = list(map(int, input().split()))
connect = floyd_warshal(graph, n)
for i in range(len(race)-1):
    if connect[race[i]-1][race[i+1]-1] != 1:
        print('NO')
        sys.exit()
print('YES')