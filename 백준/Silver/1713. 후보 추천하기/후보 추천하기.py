import sys 
input = sys.stdin.readline
sys.setrecursionlimit(10**6)

n = int(input())
r= int(input())
candidate = {}
person = list(input().split())
for p in person:
    if len(candidate)<n:
        if p not in candidate:
            candidate[p]=1
        else:
            candidate[p]+=1
    else:
        if p not in candidate:
            find = min(candidate.values())
            res = [k for k, v in candidate.items() if v==find]
            del candidate[res[0]]
            candidate[p]=1
        else:
            candidate[p]+=1
for k in sorted(list(map(int, candidate.keys()))):
    print(k, end=' ')