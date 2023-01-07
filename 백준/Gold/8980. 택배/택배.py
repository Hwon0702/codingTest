import sys 
input = sys.stdin.readline
sys.setrecursionlimit(10**6)

n, c = map(int, input().split())
m = int(input())
boxs = []
for _ in range(m):
    s, e, b = map(int, input().split())
    boxs.append([s, e, b])
boxs.sort(key=lambda x:x[1]) #도착마을 순 정렬
cost = [c]*(n+1)
total = 0
for i in range(m):
    temp = c  
    for j in range(boxs[i][0], boxs[i][1]):
        temp = min(temp, cost[j])
    temp = min(temp, boxs[i][2])
    for j in range(boxs[i][0], boxs[i][1]):
        cost[j] -= temp
    total += temp
print(total)