from collections import deque
import sys 
input = sys.stdin.readline
sys.setrecursionlimit(10**6)

graph = []
while True:
    try:
        graph.append(int(input()))
    except:
        break

def post(s, e):
    if s>e:
        return
    mid = e+1
    for i in range(s+1, e+1):
        if graph[s]<graph[i]:
            mid = i
            break
    post(s+1, mid-1)
    post(mid, e)
    print(graph[s])

post(0, len(graph)-1)