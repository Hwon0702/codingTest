import sys 
input = sys.stdin.readline
sys.setrecursionlimit
n = int(input())
ropes = [ 0 for _ in range(n)]
for i in range(n):
    ropes[i]=int(input())
ropes.sort(reverse=True)
for i in range(n):
    ropes[i] = ropes[i] * (i + 1)
print(max(ropes))