import sys 
cnt = {0:1, 1:1, 2:3, 3:5}
while(True):
    try:
        n = int(sys.stdin.readline())
        if n in cnt:
            print(cnt[n])
        else:
            for i in range(3, n+1):
                if n in cnt:
                    continue
                else:
                    cnt[i]=cnt[i-2]*2+cnt[i-1]
            print(cnt[n])
    except:
        break