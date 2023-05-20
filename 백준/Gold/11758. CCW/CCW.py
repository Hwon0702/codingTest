import sys 
input = sys.stdin.readline 
sys.setrecursionlimit(10**6)
def ccw(p1,p2,p3):
    res= (p1[0]*p2[1] + p2[0]*p3[1] + p3[0]*p1[1] - (p2[0]*p1[1] + p3[0]*p2[1] + p1[0]*p3[1]))
    if res > 0:
        return 1
    elif res == 0:
        return 0
    else:
        return -1   

points = []
for _ in range(3):
    points.append(list(map(int, input().split())))

print(ccw(points[0],points[1],points[2]))
