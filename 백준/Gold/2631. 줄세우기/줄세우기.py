import sys 
input = sys.stdin.readline
sys.setrecursionlimit(10**6)

n = int(input())
res = [1 for _ in range(n)]
numbers = []
for _ in range(n):
    numbers.append(int(input()))
for i in range(n):
    for j in range(i):
        if numbers[j]<numbers[i]:
            res[i]=max(res[i],res[j]+1)
print(n-max(res))
