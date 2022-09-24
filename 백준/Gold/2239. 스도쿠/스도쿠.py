import sys 
input = sys.stdin.readline
def rowSearch(row_num, target_num):
    for c in range(9):
        if target[row_num][c] == target_num: #가로에 넣을 숫자가 이미 있음->False
            return False
    return True
def columnSearch(col_num, target_num):
    for r in range(9):
        if target[r][col_num] == target_num: #세로에 넣을 숫자가 이미 있음->False
            return False
    return True
def squareSearch(row_num, col_num, target_num): #사각형 안에 넣을 숫자가 이미 있음->False
    nr = (row_num//3) * 3
    nc = (col_num//3) * 3
    for i in range(3):
        for j in range(3):
            if target[nr+i][nc+j] == target_num:
                return False
    return True
def dfs(depth):
    if depth == len(zeros): #0끝까지 다 찾음
        for row in range(9):
            for col in range(9):
                print(target[row][col], end="")
            print()
        exit()
    
    r, c = zeros[depth]	# 0의 위치
    for num in range(1, 10):
        if rowSearch(r, num) and columnSearch(c, num) and squareSearch(r, c, num):#세 곳 모두에 없음
            target[r][c] = num
            dfs(depth + 1)
            target[r][c] = 0


target = [list(map(int, input().rstrip()))for _ in range(9)]
zeros = [(i,j) for i in range(9) for j in range(9) if not target[i][j]]
dfs(0)