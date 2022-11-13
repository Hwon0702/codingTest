import sys 
input = sys.stdin.readline
n = int(input())
time, profit = [0 for _ in range(n+1)], [0 for _ in range(n+1)]
for i in range(n):
    t,p = map(int, input().split())
    time[i] = t
    profit[i] = p

max_profit =[0 for _ in range(n+1)]

for i in range(len(time)-2, -1, -1):  
    if time[i]+i <= n:      
        max_profit[i] = max(profit[i] + max_profit[i + time[i]], max_profit[i+1])   
    else:               
        max_profit[i] = max_profit[i+1]
print(max_profit[0])
