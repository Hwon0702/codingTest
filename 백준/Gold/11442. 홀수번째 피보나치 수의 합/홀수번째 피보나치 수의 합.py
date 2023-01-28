import sys 
input = sys.stdin.readline
sys.setrecursionlimit(10**6)

matrix = [[1, 1], [1, 0]]

def mul_matrix(matrix1, matrix2):
    res = [[0, 0] for _ in range(2)]
    for i in range(2):
        for j in range(2):
            for k in range(2):
                res[i][j] += matrix1[i][k] * matrix2[k][j]% 1000000007
    return res

def power(a, b):
    if b == 1:
        return a
    else:
        tmp = power(a, b // 2) 
        if b % 2 == 0:
            return mul_matrix(tmp, tmp) 
        else:
            return mul_matrix(mul_matrix(tmp, tmp), a) 

n = int(input())
if n%2!=0:
    n = n+1
result = power(matrix, n)
print((result[0][1])% 1000000007)
