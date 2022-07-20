import sys 
input = sys.stdin.readline
sys.setrecursionlimit(10**6)
n = int(input())
numbers = list(map(int, input().split()))
res = [1 for _ in range(n)]
for i in range(1, n):
    for j in range(i):
        if numbers[i]>numbers[j]:
            res[i] = max(res[i], res[j]+1)
print(max(res))