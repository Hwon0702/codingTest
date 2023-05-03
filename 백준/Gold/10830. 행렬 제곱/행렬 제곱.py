import sys 
input = sys.stdin.readline 
sys.setrecursionlimit(10**6)
n, m = map(int, input().split())
matrix = []
def multiple(mat1, mat2):
    res = [[0]*n for _ in range(n)]
    for i in range(n):
        for j in range(n):
            for z in range(n):
                res[i][j] += mat1[i][z] * mat2[z][j] % 1000
    return res

def power(a, b): #쪼개서 구하기
    if b == 1: 
        return a
    else:
        div = power(a, b // 2)
        if b % 2 == 0:
            return multiple(div, div)  
        else:
            return multiple(multiple(div, div), a) 

for _ in range(n):
    matrix.append(list(map(int, input().split())))
result = power(matrix, m)
for row in result:
    for col in row:
        print(col % 1000, end=' ') 
    print()