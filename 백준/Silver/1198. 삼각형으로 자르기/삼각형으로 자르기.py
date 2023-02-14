import sys 
input = sys.stdin.readline 
sys.setrecursionlimit(10**6)

n = int(input())
points_x = []
points_y = []
for _ in range(n):
    a, b = map(int, input().split())
    points_x.append(a)
    points_y.append(b)
ans = 0
for i in range(n):
    for j in range(i + 1, n):
        for k in range(j + 1, n): 
            ans = max(ans, abs(points_x[i] * points_y[j] + points_x[j] * points_y[k] + points_x[k] * points_y[i] - points_x[j] * points_y[i] - points_x[k] * points_y[j] - points_x[i] * points_y[k]));
print(ans/2)
