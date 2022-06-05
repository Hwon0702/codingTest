import sys
input = sys.stdin.readline
sys.setrecursionlimit(10**6)
n = int(input())
points = [0]*(n+1)
for i in range(1, n+1):
    points[i] = int(input())
results = [0]*(n+1)

if n==1:
    print(points[1])#마지막은 무조건 올라야함
elif n==2:
    print(points[1]+points[2]) #하나보단 당연히 두 개 올라가는게 더 큼. 포인트가 음수라는 조건 X
else:
    results[1] = points[1]
    results [2] = points[1] + points[2]
    results [3] = max(points[1] + points[3] ,points[2] + points[3])
    for i in range(4, n+1):
        results[i] = max( results[i-3] + points[i-1] + points[i] ,  results[i-2] + points[i] ) 
    print(results[n]) #마지막