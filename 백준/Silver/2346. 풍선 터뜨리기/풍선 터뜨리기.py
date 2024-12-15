import sys 
input = sys.stdin.readline
sys.setrecursionlimit(10**6)

from collections import deque
n = int(input())
q = deque(enumerate(map(int,input().split())))
nodes=[]

while q:
    idx,num = q.popleft()
    nodes.append(idx+1)
    if num>0:
        q.rotate(-(num-1))
    elif num<0:
        q.rotate(-num)
print(' '.join(map(str,nodes)))