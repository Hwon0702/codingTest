import sys 
import heapq
input = sys.stdin.readline
sys.setrecursionlimit(10**6)
n = int(input())
q = []
for i in range(n):
	num_list = list(map(int, input().split()))
	if not q: 
		for num in num_list:
			heapq.heappush(q, num)
	else:
		for num in num_list:
			if q[0] < num:
				heapq.heappush(q, num)
				heapq.heappop(q)
print(q[0])