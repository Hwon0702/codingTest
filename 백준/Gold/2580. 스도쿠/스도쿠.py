import sys 
input = sys.stdin.readline
def rowSearch(row_num, sdoku_num):
    for c in range(9):
        if sdoku[row_num][c] == sdoku_num: #가로에 넣을 숫자가 이미 있음->False
            return False
    return True
def columnSearch(col_num, sdoku_num):
    for r in range(9):
        if sdoku[r][col_num] == sdoku_num: #세로에 넣을 숫자가 이미 있음->False
            return False
    return True
def squareSearch(row_num, col_num, sdoku_num): #사각형 안에 넣을 숫자가 이미 있음->False
    nr = (row_num//3) * 3
    nc = (col_num//3) * 3
    for i in range(3):
        for j in range(3):
            if sdoku[nr+i][nc+j] == sdoku_num:
                return False
    return True
def dfs(depth):
    if depth == len(zeros): #0끝까지 다 찾음
        for row in range(9):
            for col in range(9):
                print(sdoku[row][col], end=" ")
            print()
        exit()
    
    r, c = zeros[depth]	# 0의 위치
    for num in range(1, 10):
        if rowSearch(r, num) and columnSearch(c, num) and squareSearch(r, c, num):#세 곳 모두에 없음
            sdoku[r][c] = num
            dfs(depth + 1)
            sdoku[r][c] = 0

sdoku = []
zeros=[]
for i in range(9):
    tmp = list(map(int, input().split()))
    for j in range(9):
        if tmp[j]==0:
            zeros.append([i,j])
    sdoku.append(tmp)

dfs(0)