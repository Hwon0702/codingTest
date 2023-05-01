import sys 
input = sys.stdin.readline 
sys.setrecursionlimit(10**6)
single = ['  *  ',' * * ','*****']
def star(l):
    if l == 3:
        return single

    arr = star(l//2)
    stars = []
    for i in arr:
        stars.append(' '*(l//2)+i+' '*(l//2))

    for i in arr:
        stars.append(i + ' ' + i)

    return stars

n = int(input())
for v in star(n):
    print(v)
