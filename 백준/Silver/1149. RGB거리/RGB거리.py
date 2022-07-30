import sys 
input = sys.stdin.readline
sys.setrecursionlimit(10**6)
n = int(input())
selected = [[0]*3 for _ in range(n)]
R=0
G=1
B=2
for i in range(n):
    r,g,b = map(int, input().split())
    if i==0:
        selected[0][R], selected[0][G], selected[0][B] =r,g,b
    else:
        selected[i][R] = min(selected[i-1][G], selected[i-1][B])+r  #R
        selected[i][G] = min(selected[i-1][R], selected[i-1][B])+g  #G
        selected[i][B] = min(selected[i-1][R], selected[i-1][G])+b  #B

print(min(selected[n-1]))