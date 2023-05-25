import sys 
input = sys.stdin.readline 
sys.setrecursionlimit(10**6)

n = int(input())
arr = []

for _ in range(n):
    name,k,e,m = list(map(str,input().split()))
    arr.append( [name, int(k),int(e),int(m)])

arr.sort(key = lambda x : (-x[1] , x[2],-x[3],x[0]) )

for v in arr :
    print(v[0])