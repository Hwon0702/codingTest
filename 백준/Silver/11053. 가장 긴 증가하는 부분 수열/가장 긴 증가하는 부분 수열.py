import sys 
input = sys.stdin.readline
sys.setrecursionlimit(10**6)
N = int(input())
res = [0 for _ in range(N)]
numbers = list(map(int, input().split()))
for i in range(N):
    for j in range(i):
        if numbers[i]>numbers[j] and res[i]<res[j]:
            res[i]=res[j]
    res[i]+=1
print(max(res))
