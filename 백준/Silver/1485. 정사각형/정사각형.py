import sys 
import math
input = sys.stdin.readline 
sys.setrecursionlimit(10**6)

def len(a, b):
    return math.sqrt((a[0]-b[0])**2 +(a[1]-b[1])**2 )


def isRect(points):
    points.sort()
    l1 = len(points[0], points[1])
    l2 = len(points[0], points[2])
    l3 = len(points[2], points[3])
    l4 = len(points[1], points[3])
    if l1==l2 and l1 ==l3 and l1 ==l4:
          if len(points[0], points[3])==len(points[1], points[2]):
                return 1
          else:
                return 0
    return 0
tc = int(input())
for _ in range(tc):
    points = []
    for i in range(4):
        a, b = map(int, input().split())
        points.append([a, b])
    print(isRect(points))