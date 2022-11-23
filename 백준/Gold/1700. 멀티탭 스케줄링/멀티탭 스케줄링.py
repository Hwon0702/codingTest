import sys 
input = sys.stdin.readline
sys.setrecursionlimit(10**6)
n, k = map(int, input().split())
tools = list(map(int, input().split()))
plug = []
cnt = 0
for i in range(len(tools)):
    if tools[i] in plug: #이미 꽂혀있음
        continue
    if len(plug)<n: #빈공간 있음
        plug.append(tools[i]) #꽂고 지나감
        continue 
    
    remove_target=0 #없앨거
    far=0#가장 먼거 
    for p in plug:
        if p not in tools[i:]: #앞으로 사용 안 함
            remove_target = p
            break
        elif tools[i:].index(p) > far:#가장 나중에 사용함
            far = tools[i:].index(p)
            remove_target = p 
    plug[plug.index(remove_target)]=tools[i] 
    cnt+=1

print(cnt)