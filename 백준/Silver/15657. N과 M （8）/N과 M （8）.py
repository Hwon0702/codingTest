import sys 
from itertools import combinations_with_replacement #중복 조합

input = sys.stdin.readline 
sys.setrecursionlimit(10**6)

n, m = map(int, input().split())

nums = list(map(int, input().split()))
nums.sort()
for numbers in list(combinations_with_replacement(nums, m)):
    for n in numbers:
        print(n, end=' ')
    print()