import sys
input = sys.stdin.readline
cnt = 0

def check(x): #퀸 놓을 수 있는지 확인
    for i in range(x):
        if visited[x] == visited[i] or abs(visited[x] - visited[i]) == x - i:
            return False
    return True


def queen(x, n):
    global cnt
    if x == n:
        cnt += 1
    else:
        for i in range(n):
            visited[x] = i
            if check(x):
                queen(x + 1, n)

n = int(input())
visited = [0 for _ in range(n)]
queen(0, n)
print(cnt)