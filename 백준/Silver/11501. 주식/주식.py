import sys 
input = sys.stdin.readline
sys.setrecursionlimit(10**6)

def find(N, stock):
    stock.reverse()
    max = stock[0]
    sum = 0
    for i in range(1, N):
        if max < stock[i]:
            max = stock[i]
        else:
            sum += max - stock[i]
    return sum    

T = int(input())
for _ in range(T):
    N=int(input())
    stock = list(map(int, input().split()))
    print(find(N, stock))