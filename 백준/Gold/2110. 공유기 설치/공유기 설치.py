import sys 
input = sys.stdin.readline
sys.setrecursionlimit(10**6)
n, c = map(int, sys.stdin.readline().split())
distance = [int(sys.stdin.readline()) for _ in range(n)]
 
distance.sort()
start = 1 
end = distance[n-1] - distance[0]
result = []
 
while start <= end:
    count = 1
    mid = (start + end) // 2
    current = distance[0] 
    for x in distance:
        if current + mid <= x: #
            count += 1
            current = x 
    if count >= c:
        start = mid + 1 
        result.append(mid)
    else:
        end = mid - 1 
print(max(result))