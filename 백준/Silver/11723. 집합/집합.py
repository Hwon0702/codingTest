import sys
input = sys.stdin.readline
sys.setrecursionlimit(10**6)

tc = int(input())
s = [0]*22
for _ in range(tc):
    methods = list(input().strip('\n').split(' '))
    m =methods[0]
    if len(methods)>1:
        n=int(methods[1])+1
        if m=='add':
            s[n]=1
        elif m=='remove':
            s[n]=0
        elif m=='check':
            print(s[n])
        elif m=='toggle':
            if s[n]>0:
                s[n]=0
            else:
                s[n]=1
    else:
        if m=='all':
            s = [1]*22
        elif m=='empty':
            s = [0]*22