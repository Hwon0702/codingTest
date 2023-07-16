import sys 
input = sys.stdin.readline 
sys.setrecursionlimit(10**6)
n, m=map(int, input().split())
order = [i for i in range(1, n+1)]
for _ in range(m):
    i, j = map(int, input().split())
    r = j - i + 1
    for _ in range(r//2):
        temp = order[i-1]
        order[i-1] = order[j-1]
        order[j-1] = temp
        i += 1
        j -= 1

for v in order:
    print(v, end=' ')