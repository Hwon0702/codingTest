import sys 
input = sys.stdin.readline 
sys.setrecursionlimit(10**6)

n, l = map(int, input().split())
fruits = list(map(int, input().split()))
fruits.sort()
for fruit in fruits:
    if fruit<=l:
        l+=1
    else:
        break
print(l)