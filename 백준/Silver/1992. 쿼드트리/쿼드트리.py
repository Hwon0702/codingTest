import sys
sys.setrecursionlimit(10**6)

def div(start_x, start_y, n):
    check = area[start_x][start_y]
    for x in range(start_x, start_x+n):
        for y in range(start_y, start_y+n):
            if check!=area[x][y]:
                check=-1
                break
    if check<0:
        print("(", end='')
        n = n//2
        div(start_x, start_y, n) #2사분면
        div(start_x, start_y+n, n) #3사분면
        div(start_x+n, start_y, n) #1사분면
        div(start_x+n, start_y+n, n) #4사분면
        print(")", end='')
    else:
        print(check, end="")
n = int(input())
area = [list(map(int, input())) for _ in range(n)]
div(0, 0, n)