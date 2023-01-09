import sys 
input = sys.stdin.readline
sys.setrecursionlimit(10**6)
def cal(a,b):
    return a**2 - b**2
g = int(input().strip('\n'))
l, r = 1,1
not_answer = True

while l+r <= g:
    v=cal(l,r) 
    if v == g :
        print(l)
        l +=1
        not_answer = False
    elif v > g :
        r +=1
    elif v < g :
        l +=1

if not_answer :
    print(-1)