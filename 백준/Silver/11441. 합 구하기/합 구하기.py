import sys
input = sys.stdin.readline
sys.setrecursionlimit(10**6)
len = int(input())
numbers = list(map(int, input().split()))
numbers.insert(0,0)
numbers_sum = [0]*(len+1)
numbers_sum[1]=numbers[1]
for i in range(1, len+1):
    numbers_sum[i]=numbers[i]+numbers_sum[i-1]
tc = int(input())
for i in range(tc):
    start, end = map(int, input().split())
    print(numbers_sum[end]-numbers_sum[start-1])