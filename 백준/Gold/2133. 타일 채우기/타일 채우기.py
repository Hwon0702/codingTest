
import sys 
input = sys.stdin.readline
sys.setrecursionlimit(10**6)
n = int(input())

res = [0 for _ in range(31)]
res[0] = 1
res[2]=3
for i in range(4, n+1, 2): 
    res[i] = res[i-2] * 3 
    for j in range(0, i-2, 2): 
        res[i] += res[j] * 2

print(res[n])