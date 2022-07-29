import sys 
input = sys.stdin.readline
sys.setrecursionlimit(10**6)

n,m=map(int, input().split())
target = []
eq = 0
for i in range(n):
    target.append(input().strip('\n'))
for i in range(m):
    s = input().strip('\n')
    for v in target:
        if s == v[:len(s)]:
            eq+=1
            break
print(eq)
