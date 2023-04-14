import sys 
input = sys.stdin.readline
sys.setrecursionlimit(10**6)

k = int(input())
method = list(input().split())
visited = [False for _ in range(10)]
max_res = ""
min_res = ""

def check(i, j, k):
    if k == "<":
        return i < j
    else:
        return i > j
def dfs(depth, s):
    global max_res, min_res
    if(depth == k+1):
        if len(min_res) == 0:
            min_res = s
        else:
            max_res = s
        return

    for i in range(10):
        if not visited[i]:
            if(depth == 0 or check(s[-1], str(i), method[depth-1])):
                visited[i] = True
                dfs(depth+1, s+str(i))
                visited[i] = False
dfs(0, "")
print(max_res)
print(min_res)