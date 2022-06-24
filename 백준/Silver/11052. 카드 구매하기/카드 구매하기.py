import sys 
input = sys.stdin.readline
sys.setrecursionlimit(10**6)
n = int(input())
numbers = list(map(int, input().split()))
numbers.insert(0, 0) #계산 편하게 하려고 끼워넣음
results = [0]*(n+1)
results[1]=numbers[1]
if n==1:
    print(numbers[1])
else:
    for i in range(2, n+1):
        results[i]=numbers[i]
        for j in range(1, i):
            results[i]=max(results[i], results[i-j]+numbers[j])
    print(results[n])