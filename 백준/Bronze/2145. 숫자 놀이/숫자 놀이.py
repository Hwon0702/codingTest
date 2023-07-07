import sys 
input = sys.stdin.readline 
sys.setrecursionlimit(10**6)
while True:
    n = str(input().rstrip('\n'))
    if n=='0':
        break
    n_list = list(n)
    if len(n_list)==1:
        print(int(n_list[0]))
    else:
        cnt = 0
        while len(n_list)>1:
            cnt = 0
            for v in n_list:
                cnt+=int(v)
            n_list = list(str(cnt))
        print(cnt)

