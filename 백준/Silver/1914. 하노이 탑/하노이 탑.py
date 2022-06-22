import sys 
input = sys.stdin.readline
sys.setrecursionlimit(10**6)

def hanoi( n, froms, to, aux ):
    if n == 1:
        print(froms, to)
    else:
        hanoi( n-1, froms, aux, to)
        print(froms, to)
        hanoi( n-1, aux, to, froms)

n = int(input())
result = 2**n-1
print(int(result))
if n<=20:
    hanoi(n, 1,3, 2)

