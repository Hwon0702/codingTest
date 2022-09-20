import sys 
input = sys.stdin.readline
sys.setrecursionlimit(10**6)
tc = int(input())

def dfs(v):
    global cnt, visited, finish
    visited[v]=True
    next = persons[v]
    if not visited[next]:
        dfs(next)
    else:
        if not finish[next]:
            i = next
            while i!=v:
                cnt+=1
                i = persons[i]
            cnt+=1
    finish[v]=True
for _ in range(tc):
    n = int(input())
    persons = [0]+list(map(int, input().split()))
    visited = [True]+[False]*(n)

    finish = [True]+[False]*(n)
    cnt = 0
    for i in range(1, n+1):
        if not visited[i]:
            dfs(i)
    print(n-cnt)
