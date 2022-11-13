import sys 
input = sys.stdin.readline
sys.setrecursionlimit(10**6)
from itertools import combinations

N,M=map(int, input().split())
numbers = list(map(int, input().split()))
numbers.sort()
combination_list = list(combinations(numbers, M))
combination_list.sort()
for v in combination_list:
    print( ' '.join(map(str,v)))