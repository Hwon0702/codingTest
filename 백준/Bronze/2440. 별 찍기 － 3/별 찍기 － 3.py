import sys 
input = sys.stdin.readline
sys.setrecursionlimit(10**6)

n = int(input())
for i in range(n):
    for j in range(n-i):
        print("*", end='')
    print()