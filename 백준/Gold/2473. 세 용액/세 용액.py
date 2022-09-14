import sys 
input = sys.stdin.readline
n = int(input())
requids = list(map(int, input().split()))
requids.sort()

result=[0, 0, 0]
min = float('inf')
for mid in range(n-2):
    start = mid+1
    end = n-1
    while start<end:
        tmp_min = requids[mid]+requids[start]+requids[end]
        if abs(tmp_min) < min:
            min = abs(tmp_min)
            result = [requids[mid],requids[start],requids[end]]
        if tmp_min < 0:
            start += 1
        elif tmp_min > 0:
            end -= 1
        else:
            break
result.sort()
print(result[0], result[1], result[2])