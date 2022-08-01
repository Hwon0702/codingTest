import sys
input = sys.stdin.readline
sys.setrecursionlimit(10**6)
X, Y = map(int, input().split())
Z = Y * 100 // X #승률
if Z >= 99: ##99가 넘어가버리면 절대 못이김
    print(-1)
else:
    answer = 0
    left = 1
    right = X
 
    while left <= right:
        mid = (left + right)//2 
        if (Y+mid)*100 // (X+mid) <= Z:
            left = mid+1
        else:
            answer = mid
            right = mid - 1
 
    print(answer)