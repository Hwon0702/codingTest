import sys 
input = sys.stdin.readline
sys.setrecursionlimit(10**6)
from itertools import permutations
n,m = map(int, input().split())
numbers = list(map(int, input().split()))
results = list(permutations(numbers, m))
res_set = set(results) 
results = list(res_set) 
results.sort()
for i in results:
    print( ' '.join(map(str,i)))