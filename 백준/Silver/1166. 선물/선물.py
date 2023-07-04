import sys 
input = sys.stdin.readline
sys.setrecursionlimit(10**6)

n, l, w, h = map(int, input().split())
left, right = 0, max(l, w, h)

for _ in range(10000):
    mid = (left + right) / 2
    count = (l // mid) * (w // mid) * (h // mid)
    if count >= n:
        left = mid
    else:
        right = mid

print("%.10f" %(right))