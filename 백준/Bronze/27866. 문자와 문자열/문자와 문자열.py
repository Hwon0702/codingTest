import sys 
input = sys.stdin.readline 
sys.setrecursionlimit(10**6)
l = list(map(str, input().rstrip('\n')))
n = int(input())
print(l[n-1])