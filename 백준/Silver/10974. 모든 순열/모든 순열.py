import sys
input = sys.stdin.readline
sys.setrecursionlimit(10**6)
from itertools import permutations
N = int(input())
numbers =  [i for i in range(1, N+1)]
results = list(permutations(numbers))
results.sort()
for v in results:
    print( ' '.join(map(str,v)))