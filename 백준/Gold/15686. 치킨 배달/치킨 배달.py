import sys 
input = sys.stdin.readline
sys.setrecursionlimit(10**6)
from itertools import combinations
n, m = map(int, input().split())
cities = []
chickens = []
for y in range(n):
    x_list = list(map(int, input().split()))
    for x in range(n):
        if x_list[x]==1:
            cities.append([x, y])
        elif x_list[x]==2:
            chickens.append([x, y])
dirs = [[float('inf') for _ in range(len(cities))] for _ in range(len(chickens))]
chicken_idx = [i for i in range(len(chickens))]
remain_chicken = []
for i in range(0, m+1):
    remain_chicken.append(list(combinations(chicken_idx, i)))
chicken_dir  =float('inf')
remain_chicken = list(combinations(chickens,m))
for remain in remain_chicken:
    city_dir =0
    for j in range(len(cities)):
        home_dir =float('inf')
        for x,y in remain:
            home_dir = min(home_dir,abs(x-cities[j][0])+abs(y-cities[j][1]))
        city_dir+=home_dir
    chicken_dir = min(chicken_dir,city_dir)

print(chicken_dir)