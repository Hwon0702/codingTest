import sys 
input = sys.stdin.readline
sys.setrecursionlimit(10**6)

    
n = int(input())
weights = list(map(int,input().split()))
num_of_marble = int(input())
list_of_marble = list(map(int,input().split()))

dp= [[0 for j in range((i+1)*500+1)] for i in range(n+1)]
tmp = []

def find(num,weight):
    if num > n:
        return 
    
    if dp[num][weight]:
        return
    
    dp[num][weight] = 1
    
    find(num+1, weight)
    find(num+1, weight+weights[num-1])
    find(num+1, abs(weight-weights[num-1]))
    
find(0,0)

for marble in list_of_marble:
    if marble > 30*500:
        tmp.append("N")
        continue
    if dp[n][marble] == 1:
        tmp.append("Y")
    else:
        tmp.append("N")
print(*tmp)