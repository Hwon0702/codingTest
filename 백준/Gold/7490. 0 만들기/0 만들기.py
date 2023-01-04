import sys 
input = sys.stdin.readline
sys.setrecursionlimit(10**6)



def search(current, before, target):
    global res
    if current ==target:
        before.append(target)
        a = ''.join(map(str,before))
        if (eval(a.replace(' ', '')))==0:
            res.append(a)
        return
    for c in cal:
        be = before[:]
        be.append(current)
        be.append(c)
        search(current+1, be, target)


cal = [' ', '+', '-']
n = int(input())
for _ in range(n):
    res = []
    num = int(input())
    search(1, [], num)
    for v in res:
        print(v)
    print()
