import sys 
input = sys.stdin.readline 
sys.setrecursionlimit(10**6)
n = int(input())
#가장 작은 좋은 수열
#가장 작으려면 : 일단 1, 2, 3의 순으로 가는게 좋음. 1213121 1-> 1+1 -> 1 -> 바로 전거보다 +1 -> 3 ->1 ->2->1->
#계속 1213으로 갈 수는 없음 : 1213이 반복되기 때문
#1 2 1 3 1 2 1 
#같은 수는 두 번 반복 될 수 없음
#
def dfs(res, idx):
    #반복되는 수열이 있는지 확인
    for i in range(1, idx//2+1):
        if res[idx-(2*i):idx-(2*i)+i] == res[idx-(2*i)+i:]:
            return
    if idx == n:
        print(''.join(map(str, res)))
        sys.exit()

    for num in (1, 2, 3):
        if res and res[-1] == num:
            continue
        res.append(num)
        dfs(res, idx+1)
        res.pop()
dfs([], 0)