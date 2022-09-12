import sys
input = sys.stdin.readline
sys.setrecursionlimit(10**6)
def divide(x, y, n):
    global blue, white 
    start = paper[x][y]
    for i in range(x, x+n):
        for j in range(y, y+n):
            if start!=paper[i][j]:
                divide(x, y, n//2)
                divide(x, y+n//2, n//2)
                divide(x+n//2, y, n//2)
                divide(x+n//2, y+n//2, n//2)
                return
    if start==0:
        white+=1
        return
    else:
        blue+=1
        return

m = int(input())
paper = [[]]*m
for i in range(m):
    paper[i] = list(map(int, input().split()))
white, blue = 0,0
divide(0,0, m)

print(white)
print(blue)
