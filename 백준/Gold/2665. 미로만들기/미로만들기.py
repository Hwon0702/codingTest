import sys
input = sys.stdin.readline
sys.setrecursionlimit(10**6)
import heapq
n = int(input())
room = []
for _ in range(n):
    room.append(list(map(int, input().rstrip())))
visited = [[0 for _ in range(n)] for _ in range(n)]
dx = [-1, 1, 0, 0]
dy = [0, 0, -1, 1]

def dijkstra(start):
    q = []
    heapq.heappush(q, [0, start[0], start[1]])
    visited[0][0] = 1
    while q:
        cnt, cx, cy = heapq.heappop(q)
        if cx == n - 1 and cy == n - 1:
            return cnt
        for i in range(4):
            nx = cx + dx[i]
            ny = cy + dy[i]
            if 0 <= nx < n and 0 <= ny < n and visited[nx][ny] == 0:
                visited[nx][ny] = 1
                if room[nx][ny] == 0:
                    heapq.heappush(q, [cnt + 1, nx, ny])
                else:
                    heapq.heappush(q, [cnt, nx, ny])

print(dijkstra([0, 0]))