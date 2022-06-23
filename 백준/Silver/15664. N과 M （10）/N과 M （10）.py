import sys 
input = sys.stdin.readline
sys.setrecursionlimit(10**6)
from itertools import combinations
n, m = map(int, input().split())
numbers = list(map(int, input().split()))
numbers.sort()
comb = list(combinations(numbers, m))
set_comb = set(comb)
comb = list(set_comb)
comb.sort()
for v in comb:
    print(' '.join(map(str,v)))