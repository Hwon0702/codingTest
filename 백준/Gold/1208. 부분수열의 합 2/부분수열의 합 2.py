import sys 
input = sys.stdin.readline
sys.setrecursionlimit(10**6)

n, s = map(int, input().split())
nums = list(map(int, input().split()))
res = 0

left = {}
def find(idx, end, add, side_left):
    global n, s, res 
    if idx==end:
        if side_left:
            if not add in left:
                left[add]=1
            else:
                left[add]+=1
        elif s-add in left:
            res+=left[s-add]
        return
    find(idx+1, end, add+nums[idx], side_left)
    find(idx+1, end, add, side_left)
find(0, n//2, 0, True)
find(n//2, n, 0, False)
if s==0: #s가 0인 경우 아무것도 없을 때도 0이어서 이 경우의 수 하나 빼줘야댐
    res-=1
print(res)