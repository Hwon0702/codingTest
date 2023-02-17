import sys 
input = sys.stdin.readline 
sys.setrecursionlimit(10**6)
n = int(input())
num = []
for _ in range(n):
    num.append(int(input().strip('\n')))
num.sort(reverse=True)
for v in num:
    print(v)