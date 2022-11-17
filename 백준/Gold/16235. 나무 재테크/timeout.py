from collections import deque
import heapq
import sys 
input = sys.stdin.readline
sys.setrecursionlimit(10**6)
#번식용
dx = [-1,0,1,-1,1,-1,0,1]
dy = [-1,-1,-1,0,0,1,1,1]
nutriment = []#양분
area = []#땅
died = deque()#죽은나무
spread = deque() #번식대상
tree = []
n, m, year = map(int, input().split())
for _ in range(n):
    nutriment.append(list(map(int, input().split())))

#tree=[[[]for _ in range(n)]for _ in range(n)] #나무 설정
area = [[5 for _ in range(n)]for _ in range(n)]#초기 땅
#print(tree)
for _ in range(m): #나무 정보
    i, j, k = map(int, input().split())
    heapq.heappush(tree, [k, i-1, j-1])
def spring():
    global tree
    global died
    global area
    global spread
    tmp_tree = []
    
    while tree:
        k, i, j = heapq.heappop(tree)
        if area[i][j]>=k:
            area[i][j]-=k 
            heapq.heappush(tmp_tree, [k+1, i, j])
            if (k+1)%5==0:
                spread.append([i, j])
        else:
            died.append([i, j, k//2])
    tree = tmp_tree

    return
def summer():
    global tree
    global died
    global area
    global spread
    
    while died:
        x, y, a = died.popleft()
        area[x][y]+=a
    return
def autumn():
    global tree
    global died
    global area
    global spread
    while spread:
        cx, cy = spread.popleft()
        for i in range(8):
            nx = cx+dx[i]
            ny = cy+dy[i]
            if 0<=nx<n and 0<=ny<n:
                heapq.heappush(tree, [1, nx, ny])
    return 
def winter():
    global tree
    global died
    global area
    global spread
    for i in range(n):
        for j in range(n):
            area[i][j]+=nutriment[i][j]
    return
for _ in range(year):
    spring()
    summer()
    autumn()
    winter()
print(len(tree))