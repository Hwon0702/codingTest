import sys 
input = sys.stdin.readline
sys.setrecursionlimit(10**6)
n=list((input().strip('\n')))
n = sorted(n, reverse=True)
sum = 0
if '0' not in n:
    print(-1)
else:
    for i in n:
        sum += int(i)
    if sum % 3 != 0 : #3의 배수
        print(-1)
    else :
        print(''.join(n))