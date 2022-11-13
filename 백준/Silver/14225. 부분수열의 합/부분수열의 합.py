import sys 
from collections import Counter
input = sys.stdin.readline
from itertools import combinations
from collections import Counter
n = int(input())
numbers = list(map(int, input().split()))
sum_arr = []
for i in range(1, n+1):
    for v in list(combinations(numbers, i)):
        sum_arr.append(sum(list(v)))
sum_arr.sort()
check_num = [i for i in range(1, sum_arr[-1]+2)]
result = Counter(check_num) - Counter(sum_arr)
print(list(result.keys())[0])