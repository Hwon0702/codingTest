import sys 
from itertools import combinations
input = sys.stdin.readline
sys.setrecursionlimit(10**6)
while True:
    numbers = list(map(int, input().split()))
    if numbers[0]==0:
        break
    lotto  = numbers[1:]
    lotto_numbers = list(combinations(lotto, 6))
    lotto_numbers.sort()
    for v in lotto_numbers:
         print( ' '.join(map(str,v)))
    print()