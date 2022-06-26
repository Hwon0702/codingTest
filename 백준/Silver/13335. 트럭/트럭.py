import sys
input = sys.stdin.readline
sys.setrecursionlimit(10**6)

n, w, l = map(int, input().split())
cars = list(map(int, input().split()))

moving_weight = []
moving_dist = []
sum = 0
count = 1
moving_weight.append(cars.pop(0))
sum+=moving_weight[0]
moving_dist.append(w)
while len(moving_dist)>0:
    count+=1
    for i in range(0,len(moving_dist)):
        moving_dist[i]=moving_dist[i]-1
    if len(moving_dist)>0 and moving_dist[0]<=0:
        _=moving_dist.pop(0)
        sum=sum-moving_weight.pop(0)
    if (sum<l or len(moving_weight)<w) and len(cars)>0:
        if cars[0]+sum>l or len(moving_weight)+1 > w:
            continue
        else:
            sum+=cars[0]
            moving_weight.append(cars.pop(0))
            moving_dist.append(w)
    
print(count)
