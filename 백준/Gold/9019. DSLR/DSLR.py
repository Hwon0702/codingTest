from collections import deque
import sys 
input = sys.stdin.readline

def D(n):
    r = n*2
    if r>9999:
        return r%10000
    else:
        return r
def S(n):
    if n==0:
        return 9999
    else:
        return n-1
def L(n):
    return (10*n+(n//1000))%10000
def R(n):
    return (n//10+(n%10)*1000)%10000
tc = int(input())
for _ in range(tc):
    a, b = map(int, input().split())
    visit = [False] * 10000
    q = deque()
    q.append((a, ""))
    cnt = 0
    while q:
        c, m = q.popleft()
        visit[c]==True
        if b==c:
            print(m)
            break
        else:
            v =D(c)
            if not visit[v]:
                q.append((v, m+"D"))
                visit[v]=True
            v= S(c)
            if not visit[v]:
                visit[v]=True
                q.append((v,m+"S"))
            v= L(c)
            if not visit[v]:
                visit[v]=True
                q.append((v,m+"L"))
            v=R(c)
            if not visit[v]:
                visit[v]=True
                q.append((v, m+"R"))
