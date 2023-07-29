import sys 
input = sys.stdin.readline
sys.setrecursionlimit(10**6)
while True:
    try:
        r = input().rstrip('\n')
        if len(r)==0:
            break
        else:
            print(r)
    except:
        break