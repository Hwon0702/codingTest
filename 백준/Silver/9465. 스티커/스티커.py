import sys 
input = sys.stdin.readline
tc = int(input())
for _ in range(tc):
    width = int(input())
    stickers = []
    for _ in range(2):
        stickers.append(list(map(int, input().split())))
    if width<=1:
        print(max(stickers[0][0], stickers[1][0]))
    else:
        res = stickers
        res[0][1]+=res[1][0]
        res[1][1]+=res[0][0]
        for i in range(2, width):
            res[0][i] = max(res[1][i-1] + stickers[0][i], res[1][i-2] + stickers[0][i])
            res[1][i] = max(res[0][i-1] + stickers[1][i], res[0][i-2] + stickers[1][i])
        print(max(res[0][width-1], res[1][width-1]))