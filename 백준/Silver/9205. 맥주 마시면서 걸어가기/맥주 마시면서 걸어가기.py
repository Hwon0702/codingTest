from collections import deque
import sys 
input = sys.stdin.readline
def bfs():
    queue = deque()
    queue.append([home[0], home[1]])
    while queue:
        x, y = queue.popleft()
        if abs(x - rock[0]) + abs(y - rock[1]) <= 1000: #맨하탄거리 / 맥주20병*50m
            return "happy"
        for i in range(store_cnt):
            if not visited[i]: #n번째 편의점 방문X
                nx, ny = store[i]
                if abs(x - nx) + abs(y - ny) <= 1000: #맥주 소진 전까지 다음 편의점 방문 가능
                    queue.append([nx, ny])
                    visited[i] = True
    return "sad"
tc = int(input())
for _ in range(tc):
    store_cnt = int(input())
    home = list(map(int, input().split()))
    store=[list(map(int, input().split())) for _ in range(store_cnt)]
    rock = list(map(int, input().split()))
    visited = [False for _ in range(store_cnt+1)] #home 제외
    print(bfs())
