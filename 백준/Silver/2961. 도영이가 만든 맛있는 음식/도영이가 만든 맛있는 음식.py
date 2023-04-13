import sys
from itertools import combinations
sys.setrecursionlimit(10**6)
input = sys.stdin.readline

n = int(input())
sour = []
bitter = []
for _ in range(n):
    s, b = map(int, input().split())
    sour.append(s)
    bitter.append(b)

diff = float('inf')

for i in range(1, n+1): 
    combs = list(combinations(list(range(n)), i))
    for comb in combs:
        s = 1
        b = 0
        for j in range(i):
            s *= sour[comb[j]]
            b += bitter[comb[j]]
        if abs(s - b) < diff:
            diff = abs(s - b)

print(diff)
