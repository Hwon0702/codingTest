import sys 
input = sys.stdin.readline
sys.setrecursionlimit(10**6)

m, n = map(int, input().split())
land = []
res = [[0 for _ in range(n+1)] for _ in range(m+1)]
for _ in range(m):
    land.append(list(map(int, input().split())))
max_length = 0

for i in range(1,m+1):
    for j in range(1,n+1):
        if land[i-1][j-1]==0: #장애물이 없음. 하나씩 증가시켜가면서 확인 가능
            res[i][j] = min(res[i-1][j],res[i][j-1],res[i-1][j-1])+1
            max_length = max(max_length, res[i][j])
print(max_length)