import sys 
input = sys.stdin.readline 
sys.setrecursionlimit(10**6)
n = int(input())
books = {}
for _ in range(n):
    s = str(input().strip('\n'))
    if s in books:
        books[s]+=1
    else:
        books[s]=1
book = ''
cnt = 0
for k, v in books.items():
    if v>cnt:
        book = k 
        cnt = v
    elif v==cnt:
        book = min(book, k) 
print(book)
