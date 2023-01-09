import sys 
input = sys.stdin.readline
sys.setrecursionlimit(10**6)
def cal(a,b):
    return a**2 - b**2
g = int(input().strip('\n'))
l, r = 1,1
not_answer = True

while l+r <= g:
    if cal(l,r) == g :
        print(l)
        l +=1
        not_answer = False
    elif cal(l,r) > g :
        r +=1
    elif cal(l,r) < g :
        l +=1

if not_answer :
    print(-1)