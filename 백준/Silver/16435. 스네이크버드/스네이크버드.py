import sys 
import heapq
input = sys.stdin.readline 
sys.setrecursionlimit(10**6)

n, l = map(int, input().split())
fruits = list(map(int, input().split()))
heapq.heapify(fruits)
while True and fruits:
    fruit = heapq.heappop(fruits)
    if fruit<=l:
        l+=1
    else:
        break
print(l)