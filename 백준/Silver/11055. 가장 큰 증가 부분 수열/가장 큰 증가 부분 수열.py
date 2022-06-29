import re
import sys 
input = sys.stdin.readline
sys.setrecursionlimit(10**6)
n = int(input())
numbers = list(map(int, input().split()))
res=[0]*(n)
#가장 크고 싶다면 : 합을 계속 비교하면서 늘려야함
#만약 작고 작다면 -> 제낌 만약 크고 크다면 -> 넣음
#ex. 1 2 3 4 5 6 7 8 9 55 10 11 12 13 14 15 16 17 17 18 20
#55를 빼고 나머지를 넣는 것이 더 큼
#각 idx까지 증가하는 수열을 만들 경우의 합 저장
tmp_sum = 0
last_idx = 0
for i in range (n):
    for j in range (i):
        if numbers[i]>numbers[j] and res[i] < res[j] :
            res[i] = res[j]
    res[i] += numbers[i]
print(max(res))