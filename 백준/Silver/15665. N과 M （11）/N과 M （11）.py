import sys 
input = sys.stdin.readline
sys.setrecursionlimit(10**6)
from itertools import product
n, m = map(int, input().split())
numbers = list(set(map(int, input().split())))
permutations_list = product(numbers, repeat=m)
permutations_list = list(set(permutations_list))
permutations_list.sort()
for v in permutations_list:
    print(' '.join(map(str, v)))