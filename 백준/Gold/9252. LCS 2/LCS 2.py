import sys 
input = sys.stdin.readline
s1 = input().strip()
s2 = input().strip()
res = [["" for _ in range (len(s2) + 1)] for _ in range(len(s1) + 1)]
for i in range(1, len(s1)+1):
    for j in range(1, len(s2)+1):
        if s1[i - 1] == s2[j - 1]:
            res[i][j]=res[i-1][j-1]+s1[i-1]
        else:
            if len(res[i - 1][j])>len(res[i][j - 1]):
                res[i][j]=res[i - 1][j]
            else:
                res[i][j] =res[i][j - 1]
print(len(res[-1][-1]))
if len(res[-1][-1])>0:
    print(res[-1][-1])
