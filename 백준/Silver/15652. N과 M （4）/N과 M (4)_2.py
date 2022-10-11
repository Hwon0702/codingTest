import sys 
from itertools import combinations_with_replacement

input = sys.stdin.readline
sys.setrecursionlimit(10**6)

n,m= map(int, input().split())
numbers = [i for i in range(1, n+1)]
comb_list=list(combinations_with_replacement(numbers, m))
result = list(set(comb_list))
result.sort()
for v in result:
    print(' '.join(map(str, v)))