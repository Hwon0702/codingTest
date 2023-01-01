import sys 
input = sys.stdin.readline
sys.setrecursionlimit(10**6)
n = int(input())
S = []
T = ''
#dir_map = {}
#for i in range(26):
#    dir_map[chr(65+i)] = i 
for _ in range(n):
    S.append(str(input().strip('\n')))
cnt = 0
i = 0
while cnt <= n and S:
    if S[0] < S[-1]:
        T += S[0]
        S = S[1:]
    elif S[0] > S[-1]:
        T += S.pop()
    else:
        ii, jj = 0, len(S)-1
        diff = False
        while ii <= jj:
            if S[ii] < S[jj]:
                T += S[0]
                S = S[1:]
                diff = True
                break
 
            elif S[ii] > S[jj]:
                T += S.pop()
                diff = True
                break

            else:
                ii += 1
                jj -= 1
        if not diff:
            T += S[0]
            S = S[1:]
    cnt += 1
    if cnt % 80 == 0:
        T += '\n'
print(T)
