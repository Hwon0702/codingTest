import sys 
input = sys.stdin.readline
sys.setrecursionlimit(10**6)
n = int(input())
k = int(input())
#b=[]
#a=[[0 for _ in range(n)] for _ in range(n)]
#for i in range(1, n+1):
#    for j in range(1, n+1):
#        a[i-1][j-1]=i*j
#    print(a[i-1])
#    #b.append()
#1 2 3
#   2   4   6
#     3     6  9
left = 1
right = k
while left <= right:
    mid = (left + right) // 2
    
    temp = 0
    for i in range(1, n+1):
        temp += min(mid//i, n) 
    if temp >= k:
        res = mid
        right = mid - 1
    else:
        left = mid + 1
print(res)