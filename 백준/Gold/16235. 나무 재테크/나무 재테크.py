from collections import deque
import sys 
input = sys.stdin.readline
dx = [-1,0,1,-1,1,-1,0,1]
dy = [-1,-1,-1,0,0,1,1,1]
nutriment = []#양분
area = []#땅

n, m, year = map(int, input().split())
tree = [[deque() for _ in range(n)]for _ in range(n)]
for _ in range(n):
    nutriment.append(list(map(int, input().split())))
area = [[5 for _ in range(n)]for _ in range(n)]#초기 땅
for _ in range(m): #나무 정보
    i, j, k = map(int, input().split())
    tree[i-1][j-1].append(k)

def ss_season():
    for i in range(n):
        for j in range(n):
            for k in range(len(tree[i][j])):
                if area[i][j]>=tree[i][j][k]:
                    area[i][j]-=tree[i][j][k]
                    tree[i][j][k]+=1
                else:
                    for _ in range(k, len(tree[i][j])):
                        area[i][j] += tree[i][j].pop() // 2
                    break

def fw_season():
    for i in range(n):
        for j in range(n):
            for k in range(len(tree[i][j])):
                if tree[i][j][k] % 5 == 0:
                    for idx in range(8):
                        nx = i + dx[idx]
                        ny = j + dy[idx]
                        if 0 <= nx < n and 0 <= ny < n:
                            tree[nx][ny].appendleft(1)
            area[i][j] += nutriment[i][j]
for _ in range(year):
    ss_season()
    fw_season()
res = 0
for i in range(n):
    for j in range(n):
        res += len(tree[i][j])
print(res)