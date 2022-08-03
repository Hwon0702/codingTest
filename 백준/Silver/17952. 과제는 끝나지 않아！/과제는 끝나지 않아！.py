import sys 
input = sys.stdin.readline
sys.setrecursionlimit(10**6)

n = int(input())
s = [] #stack
get_point = 0
time, point = 0,0
for i in range(n):
    p = input().strip('\n')
    if len(p)>1: #새로운 과제 넣음
        _, np,nt = map(int, p.split())
        s.append((nt, np)) #새로 받은거 넣음
    if s:
        time, point = map(int, s.pop())
        if time-1==0:
            get_point+=point
        else:
            s.append((time-1, point))
    else:
        continue
print(get_point)    