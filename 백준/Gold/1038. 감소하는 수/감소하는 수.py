import sys 
from itertools import combinations
input = sys.stdin.readline
sys.setrecursionlimit(10**6)

n = int(input())
res = []
for i in range(1, 11): #1부터 10까지
    for j in combinations(range(10), i): #10까지의 수 중에 n개 뽑아서(중복X)
        num = sorted(list(j), reverse=True)  #내림차순으로 정렬돼있다면
        res.append(int("".join(map(str, num)))) #이걸 수로 만들어서 append
res.sort() #결과 정렬
if len(res)>n:
    print(res[n])
else:
    print(-1)
