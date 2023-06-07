import sys 
input = sys.stdin.readline 
sys.setrecursionlimit(10**6)
def inCircle(x, y,a, b):
    return ((x - a) * (x - a)) + ((b - y) * (b - y))
w, h, x, y, p = map(int, input().split())
cnt = 0
for _ in range(p):
    x1, y1 = map(int, input().split())
    r = (h / 2) ** 2
    if (x <= x1 and x1 <= x + w and y <= y1 and y1 <= y + h):
        cnt += 1
    # 왼쪽 
    elif (inCircle(x, y + (h / 2), x1, y1) <= r):
        cnt += 1
    # 오른쪽 
    elif (inCircle(x + w, y + (h / 2), x1, y1) <= r):
        cnt += 1
print(cnt)