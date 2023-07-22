import sys 
input = sys.stdin.readline 
sys.setrecursionlimit(10**6)
n=int(input())
for _ in range(n):
    s =list(input().rstrip('\n'))
    print('{}{}'.format(s[0],s[-1]))