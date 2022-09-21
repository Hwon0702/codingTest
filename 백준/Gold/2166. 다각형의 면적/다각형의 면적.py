import sys 
input = sys.stdin.readline
sys.setrecursionlimit(10**6)
n = int(input())
x_list = []
y_list = []
for _ in range(n):
    x, y =map(int, input().split())
    x_list.append(x)
    y_list.append(y)
x_list.append(x_list[0])
y_list.append(y_list[0])
sum_v = 0
for i in range(len(y_list)-1):
    sum_v+=x_list[i]*y_list[i+1]
    sum_v-=x_list[i+1]*y_list[i]
print("%.1f"% abs(sum_v/2))
