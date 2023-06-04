import sys 
input = sys.stdin.readline 
sys.setrecursionlimit(10**6)
n = int(input())
coordinates = []
for _ in range(n):
    x, y = map(int, input().split())
    coordinates.append([x, y])
coordinates.sort()

def dist(a, b):
    return (a[0]-b[0])**2+(a[1]-b[1])**2


def find(start,end):
    if start == end:
        return float('inf')
    if end - start == 1:
        return dist(coordinates[start],coordinates[end])
    mid = (start + end)//2
    mid_dist = min(find(start, mid), find(mid,end))
    target_coordinates = []
    for i in range(start, end+1):
        if (coordinates[mid][0] - coordinates[i][0])**2 < mid_dist:
            target_coordinates.append(coordinates[i])
    target_coordinates.sort(key=lambda x: x[1])
    
    t = len(target_coordinates)
    for i in range(t-1):
        for j in range(i+1,t):
            if (target_coordinates[i][1] - target_coordinates[j][1])**2 < mid_dist:
                mid_dist = min(mid_dist, dist(target_coordinates[i],target_coordinates[j]))
            else:
                break
    return mid_dist

print(find(0,n-1))