import sys 
from collections import deque
input = sys.stdin.readline 
sys.setrecursionlimit(10**6)

n = int(input())
people = list(map(int, input().split()))
stack = []
last = 1
for person in people:
    stack.append(person)
    while stack and stack[-1] == last:
        stack.pop() 
        last +=1 
if stack: 
    print('Sad') 
else:
    print('Nice')