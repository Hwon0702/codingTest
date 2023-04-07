import sys 
input = sys.stdin.readline 
sys.setrecursionlimit(10**6)
n = int(input())
height = []
width = []
total = []
for i in range(6):
    dir, length = map(int, input().split())
    if dir == 1 or dir ==2:
        width.append(length)
        total.append(length)
    else:
        height.append(length)
        total.append(length)

total_size = max(height) * max(width)
max_height = total.index(max(height))
max_width = total.index(max(width))
small_height = 0
small_width = 0
if max_height == 5:
    small_height = abs(total[max_height-1] - total[max_height-5])
else:
    small_height = abs(total[max_height-1] - total[max_height+1])
if max_width == 5:
    small_width = abs(total[max_width-1] - total[max_width-5])
else:
    small_width = abs(total[max_width-1] - total[max_width+1])
area = total_size - (small_height * small_width)
print(area*n)