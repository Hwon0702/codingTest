import sys
input = sys.stdin.readline
sys.setrecursionlimit(10**6)
def colorMatch(color, n):
    
    count =1
    drawing = [False for _ in range (n)]
    
    for i in range(n):
        if colors[i]==color:
            drawing[i]=True
    idx = -1
    for i in range(n):
        if drawing[i]==False:
            if idx<0:
                idx = i
                count+=1
            elif idx==(i-1):
                idx = i
                continue
            else:
                idx = i
                count+=1
    return count

n = int(input())
colors = list(map(str, input().strip('\n')))
b_count = colorMatch('B', n)
r_count= colorMatch('R', n)
print(min(b_count, r_count))


