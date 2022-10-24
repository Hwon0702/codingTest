import sys 
input = sys.stdin.readline
sys.setrecursionlimit(10**6)

target = int(input())
n = int(input())
cant_used = list(map(int, input().split()))

min_count = abs(100 - target)
for compare in range(1000001):
    compare = str(compare)
    for v in range(len(compare)):
        if int(compare[v]) in cant_used:
            break
        elif v == len(compare) - 1:
            min_count = min(min_count, abs(int(compare) - target) + len(compare))
print(min_count)