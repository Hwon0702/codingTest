import sys 
input = sys.stdin.readline
sys.setrecursionlimit(10**6)
tc = int(input())
for _ in range(tc):
    n, m = map(int, input().split())
    books = [False for _ in range(n+1)]
    req = []
    for _ in range(m):
        req.append(list(map(int, input().split())))
    req.sort(key=lambda x: x[1])
    cnt = 0
    for a, b in req:
        for i in range(a, b+1):
            if not books[i]:
                cnt += 1
                books[i] = True
                break
    print(cnt)