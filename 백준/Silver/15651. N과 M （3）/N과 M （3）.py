import sys
input = sys.stdin.readline
sys.setrecursionlimit(10**6)
from itertools import product
N,M=map(int, input().split())
arr =[i for i in range(1,N+1)]
all = list(product(arr, repeat=M))
all.sort()
for i in all:
    print( ' '.join(map(str,i)))
