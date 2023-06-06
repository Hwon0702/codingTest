import sys 
input = sys.stdin.readline 
sys.setrecursionlimit(10**6)
def gradient(a, b):
    return (b[1]-a[1])/(b[0]-a[0])
n = int(input())
building = list(map(int, input().split()))

res = 0
for idx, b in enumerate(building):
    cnt = 0
    left = float('inf')
    right = -float('inf') 
    for i in range(idx-1,-1,-1):  
        calc = gradient([idx+1,b],[i+1,building[i]]) 
        if calc < left:
            left = calc 
            cnt += 1
    for i in range(idx+1,n):
        calc = gradient([idx+1,b],[i+1,building[i]])
        if calc > right: 
            right = calc
            cnt += 1
    res = max(res,cnt)
print(res)