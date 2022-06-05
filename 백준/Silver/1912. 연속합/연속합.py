import sys
from unittest import result 
input = sys.stdin.readline
sys.setrecursionlimit(10**6)

n = int(input())
numbers = list(map(int, input().split()))

sum_arr = [0]*len(numbers)
sum_arr[0]=numbers[0]
for i in range(1, len(numbers)):
    sum_arr[i] = max(sum_arr[i-1]+numbers[i], numbers[i])
print(max(sum_arr))