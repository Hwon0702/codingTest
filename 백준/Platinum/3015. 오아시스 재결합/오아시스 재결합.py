import sys 
input = sys.stdin.readline
sys.setrecursionlimit(10**6)
n = int(input())
person = []
res = 0

for _ in range(n):
    now = int(input())
    cnt = 1

    while person and person[-1][0] <= now:
        height, cnt = person.pop()
        if height == now:
            res += cnt
            cnt += 1
        elif height < now: #사이에 키가 더 큰 사람이 있음
            res += cnt
            cnt = 1

    if person:
        res += 1
    person.append((now, cnt))

print(res)