import sys
input = sys.stdin.readline 
sys.setrecursionlimit(10**6)
n, m = map(int, input().split())
urlPw = dict()
for i in range(n):
    url, pw = map(str, input().split())
    urlPw[url]=pw
for i in range(m):
    print(urlPw[input().rstrip('\n')])