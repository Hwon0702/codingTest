import sys 
import heapq
input = sys.stdin.readline
sys.setrecursionlimit(10**6)
n = int(input())
room = []
q=[]
for _ in range(n):
    s, e = map(int, input().split())
    q.append([s, e])
q.sort()
heapq.heappush(room, (q[0][1]))
for i in range(1, n):
    if q[i][0]<room[0]:
        heapq.heappush(room, q[i][1])
    else:
        heapq.heappop(room)
        heapq.heappush(room, q[i][1])
print(len(room))