import sys 
input = sys.stdin.readline
sys.setrecursionlimit(10**6)
from itertools import permutations
n,m = map(int, input().split())
numbers = [i for i in range(1, n+1)]
comb = list(permutations(numbers, m))
comb.sort()
for v in comb:
    print( ' '.join(map(str,v)))