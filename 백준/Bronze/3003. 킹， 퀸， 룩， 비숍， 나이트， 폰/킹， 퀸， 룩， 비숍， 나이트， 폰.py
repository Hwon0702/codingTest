import sys 
input = sys.stdin.readline
sys.setrecursionlimit(10**6)
right = [1,1,2,2,2,8]
search = list(map(int, input().split()))
for i, v in enumerate(right):
    right[i]=v-search[i]
    print(right[i], end=' ')