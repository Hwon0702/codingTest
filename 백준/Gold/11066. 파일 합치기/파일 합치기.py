import sys 
input = sys.stdin.readline 
tc = int(input())
for _ in range(tc):
    c = int(input())
    chapters = [0]+list(map(int, input().split()))

    sums = [0 for _ in range(c+1)]

    for i in range(1,c+1):
        sums[i] = sums[i-1] + chapters[i]
    dp = [[0 for _ in range(c+1)] for _ in range(c+1)]

    for i in range(2,c+1):
        for j in range(1,c+2-i):
            dp[j][j+i-1] = min([dp[j][j+q] + dp[j+q+1][j+i-1] for q in range(i-1)]) +(sums[j+i-1] - sums[j-1])
    print(dp[1][c])