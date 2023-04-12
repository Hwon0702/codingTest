import sys 
input = sys.stdin.readline 
sys.setrecursionlimit(10**6)
dx = [-1,1,0,0]
dy = [0,0,-1,1]
res = set()

def dfs(x, y, ns):
    global res
    if len(ns)==6:
        res.add(ns)
        return
    else:
        for i in range(4):
            nx = x+dx[i]
            ny = y+dy[i]
            if 0<=nx<5 and 0<=ny<5:
                dfs(nx, ny, ns+nums[nx][ny])
nums = []
for _ in range(5):
    nums.append(list(map(str, input().split())))
for i in range(5):
    for j in range(5):
        dfs(i, j, nums[i][j])
print(len(res))