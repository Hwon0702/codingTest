import sys 
input = sys.stdin.readline
sys.setrecursionlimit(10**6)
from itertools import permutations
n,m = map(int, input().split())
numbers = []
numbers = list(map(int, input().split()))
comb = list(permutations(numbers, m))
comb.sort()
for i in comb:
    print( ' '.join(map(str,i)))