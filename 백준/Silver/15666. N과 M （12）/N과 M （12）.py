import sys
input = sys.stdin.readline
sys.setrecursionlimit(10**6)
from itertools import combinations_with_replacement

n,m= map(int, input().split())
numbers = list(map(int, input().split()))
numbers.sort()
comb_list=list(combinations_with_replacement(numbers, m))
result = list(set(comb_list))
result.sort()
for v in result:
    print(' '.join(map(str, v)))