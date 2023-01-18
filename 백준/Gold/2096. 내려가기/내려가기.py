import sys 
input= sys.stdin.readline
sys.setrecursionlimit(10**6)

n = int(input())

max_res = [0 for _ in range(3)]
min_res = [0 for _ in range(3)]
now_max = [0 for _ in range(3)]
now_min = [0 for _ in range(3)]
for _ in range(n):
    tmp = list(map(int, input().split()))
    for i in range(3):
        if i == 0:
            now_max[0] = tmp[0] + max(max_res[0], max_res[1])
            now_min[0] = tmp[0] + min(min_res[0], min_res[1])
        elif i == 1:
            now_max[1] = tmp[1] + max(max_res[0], max_res[1], max_res[2])
            now_min[1] = tmp[1] + min(min_res[0], min_res[1], min_res[2])
        else:
            now_max[2] = tmp[2]+ max(max_res[1], max_res[2])
            now_min[2] = tmp[2]+ min(min_res[1], min_res[2])

    for i in range(3):
        max_res[i] = now_max[i]
        min_res[i] = now_min[i]

print(max(max_res), min(min_res))