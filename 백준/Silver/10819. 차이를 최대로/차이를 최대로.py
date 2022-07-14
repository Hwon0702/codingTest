import sys 
input = sys.stdin.readline
from itertools import permutations
sys.setrecursionlimit(10**6)

n = int(input())
numbers = list(map(int, input().split()))
numbers_permutations = list(permutations(numbers, n))
max_result=0
for perm in numbers_permutations:
    diff_sum=0
    for i in range(1, len(perm)):
        diff_sum+=abs(perm[i-1]-perm[i])
    max_result = max(diff_sum,max_result)
print(max_result)
