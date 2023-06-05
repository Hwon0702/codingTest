import sys 
import math
input = sys.stdin.readline 
sys.setrecursionlimit(10**6)
def dist(a, b):
    return math.sqrt(abs((a[0]-b[0])**2 + (a[1]-b[1])**2))
x, y, d, t = map(int, input().split())
dist1 = dist([x, y], [0,0]) #x, y에서 바로 0,0으로 가는 길이
if t > d:
    print(dist1)
else:
    n = dist1 // d
    if d<dist1:
        print(min(t*n+dist1-d*n,t*(n+1)))
    else:
        print(min(dist1,t+d-dist1,2*t))

