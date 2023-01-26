import sys 
input = sys.stdin.readline
sys.setrecursionlimit(10**6)

n = int(input())
matrix = [[1, 1], [1, 0]]

# 행렬 곱셈
def mul_matrix(matrix1, matrix2):
    res = [[0]*2 for _ in range(2)]
    for i in range(2):
        for j in range(2):
            for k in range(2):
                res[i][j] += matrix1[i][k] * matrix2[k][j]% 1000000007
    return res

def power(a, b):
    if b == 1:  # b의 값이 1이 될 때까지 재귀
        return a
    else:
        tmp = power(a, b // 2)  # a^(b // 2)
        if b % 2 == 0:
            return mul_matrix(tmp, tmp)  # b%2==0
        else:
            return mul_matrix(mul_matrix(tmp, tmp), a)  #else
result = power(matrix, n)

print(result[0][1]% 1000000007)