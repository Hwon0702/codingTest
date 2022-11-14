from collections import deque
import sys 
input = sys.stdin.readline
sys.setrecursionlimit(10**6)
n = int(input())
numbers = list(map(int, input().split()))
visited = [False for _ in range(n)]
def move():
    q = deque()
    q.append([0, numbers[0], 0])

    while q:
        x,move, cnt = q.popleft()
        if x==n-1:
            return cnt
        for i in range(1, move+1):
            if 0<=x+i<n and visited[x+i]==False:
                visited[x+i]=True
                q.append([x+i, numbers[x+i], cnt+1])
    return -1 
print(move())