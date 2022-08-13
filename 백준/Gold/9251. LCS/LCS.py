import sys 
input = sys.stdin.readline
sys.setrecursionlimit(10**6)
s1 = input().strip('\n')
s2 = input().strip('\n')
res = [[0 for _ in range (len(s2) + 1)] for _ in range(len(s1) + 1)]
cnt = 0
for i in range(1, len(s1)+1):
    for j in range(1, len(s2)+1):
        if s1[i - 1] == s2[j - 1]:
            res[i][j]=res[i-1][j-1]+1
        else:
            res[i][j] = max(res[i - 1][j], res[i][j - 1])
print(res[-1][-1])