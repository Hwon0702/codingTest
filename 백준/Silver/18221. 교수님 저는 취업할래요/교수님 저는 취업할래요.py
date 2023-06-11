import sys 
input = sys.stdin.readline
sys.setrecursionlimit(10**6)

sk=[]
professor=[]
n = int(input())
graph = []
for i in range(n):
    tmp = list(map(int, input().split()))
    for j in range(n):
        if tmp[j]==2:
            sk = [i, j]
        elif tmp[j]==5:
            professor=[i, j]
    graph.append(tmp)
def find(x1, y1, x2, y2):
    cnt = 0
    x_min,x_max = min(x1,x2),max(x1,x2)
    y_min,y_max = min(y1,y2),max(y1,y2)
    for y in range(y_min,y_max+1):
        for x in range(x_min,x_max+1):
            if graph[y][x]==1:
                cnt+=1
    if cnt>=3 and (x1-x2)**2 + (y1-y2)**2 >=25:
        if x1==x2 or y1==y2:
            return True
        else:
            if (x1-x2)**2 + (y1-y2)**2 >=25:
                return True
    else:
        return False
        
y1,x1 = sk
y2,x2 = professor
res = find(x1,y1,x2,y2)
print(int(res))
