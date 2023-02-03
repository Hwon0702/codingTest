import sys 
input = sys.stdin.readline 
sys.setrecursionlimit(10**6)
n, m = map(int, input().split())
times = []
for _ in range(n):
    times.append(int(input()))

left,right = min(times), max(times) * m #최소 : 한 번에 모조리 끝, 최대 : 모든 사람이 최장 시간 걸림
res = right 

while left <= right:
    person = 0 
    mid = (left + right) // 2
    for i in range(n):
        person += mid // times[i] #지금 구한 시간으로 통과 가능한 사람
    if person >= m:
        right = mid - 1
        res = min(res, mid)
    else:
        left = mid + 1
print(res)