import sys 
input = sys.stdin.readline
sys.setrecursionlimit(10**6)

bar =list(input())
res = 0
s = []
for i in range(len(bar)):
    if bar[i]=='(':
        s.append(bar[i])
    elif  len(s)>0:
        if bar[i-1]=='(':
            s.pop()
            res+=len(s)
        else:
            s.pop()
            res+=1
print(res)
