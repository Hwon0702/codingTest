import sys
input = sys.stdin.readline
sys.setrecursionlimit(10**6)
len = int(input())
numbers = list(map(int, input().split()))
res = [1]*len
for i in range(len):
    for j in range(i):
        if numbers[i]<numbers[j] :
            res[i]=max(res[i], res[j]+1)
print(max(res))