import sys
input = sys.stdin.readline
sys.setrecursionlimit(10**6)


def dfs(start, cnt):
    global answer
    if answer == 1:  # 답을 이미 찾았다면 dfs 멈추기
        return

    if cnt == 5:  # 답 찾았을 때 (친구가 4명)
        answer = 1
        return

    visit[start] = 1
    for next in friends[start]:
        if not visit[next]:
            dfs(next, cnt + 1)

    visit[start] = 0


P, F = map(int, input().split())
visit = [False]*P #사람수만큼 false 
friends = [[] for _ in range(P)]
answer = 0
for i in range(F):
    p, f = map(int, input().split())
    friends[p].append(f)
    friends[f].append(p)
for i in range(P):
    if not visit[i]:
        dfs(i, 1)
print(answer)