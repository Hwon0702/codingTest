import sys 
input = sys.stdin.readline 
sys.setrecursionlimit(10**6)
n, m = map(int,input().split())
ball = [0 for _ in range(n+1)]
for i in range(m):
    i, j, k = map(int, input().split())
    ball=ball[0:i]+[k] * (j-i+1)+ball[j+1:n+1]
print(*ball[1:])
