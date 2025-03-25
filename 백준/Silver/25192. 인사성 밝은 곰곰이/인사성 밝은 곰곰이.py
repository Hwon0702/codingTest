import sys 
input = sys.stdin.readline
sys.setrecursionlimit(10**6)

n=int(input())
cnt = 0
m = dict()
for i in range(n):
    s = input().strip()
    if s=='ENTER':
        m = dict()
    else:
        if not m.get(s):
            m[s]=True
            cnt+=1
print(cnt)