import sys 
input = sys.stdin.readline
sys.setrecursionlimit(10**6)

n = int(input())
bee = 1
cnt = 1
while n > bee :
    bee += 6 * cnt 
    cnt += 1
print(cnt)