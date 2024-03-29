import sys 
input = sys.stdin.readline
sys.setrecursionlimit(10**6)
s1 = input().strip()
s2 = input().strip()
s3 = input().strip()
res = [[[0 for _ in range (len(s3) + 1)] for _ in range(len(s2) + 1)] for _ in range(len(s1) + 1)]
for i in range(1, len(s1)+1):
    for j in range(1, len(s2)+1):
        for k in range(1, len(s3)+1):
            if s1[i - 1] == s2[j - 1] == s3[k - 1]:
                res[i][j][k]=res[i-1][j-1][k-1]+1
            else:
                res[i][j][k]=max(res[i - 1][j][k],res[i][j - 1][k],res[i][j][k-1])
print(res[-1][-1][-1])