import sys 
input = sys.stdin.readline
sys.setrecursionlimit(10**6)

def find(p):
    tmp = [0] * len(p)
    
    j = 0
    for i in range(1, len(p)):
        while j > 0 and p[i] != p[j]:
            j = tmp[j-1]

        if p[i] == p[j]:
            j += 1
            tmp[i] = j

    return tmp


def KMP(s, p):
    cnt = 0
    pos = []

    j = 0
    for i in range(len(s)):
        while j > 0 and s[i] != p[j]:
            j = tb[j-1]

        if s[i] == p[j]:
            if j == len(p)-1:
                cnt += 1
                pos.append(i - len(p) + 2)
                j = tb[j]
            else:
                j += 1

    return cnt, pos


string = input().rstrip('\n')
pattern = input().rstrip('\n')
tb = find(pattern)

res, positions = KMP(string, pattern)
print(res)
print(*positions)