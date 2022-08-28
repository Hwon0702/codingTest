import sys 
input = sys.stdin.readline
sys.setrecursionlimit(10**6)
def search(idx):
    global result
    if idx == len(target): # 정답
        result = 1
        return
    if dp[idx]: # 이미 검사
        return
    dp[idx] = 1 # 검사했으니 1로 만들어 줌
    for i in range(len(compare)):
        if len(target[idx:]) >= len(compare[i]):
            for j in range(len(compare[i])): 
                if compare[i][j] != target[idx+j]: 
                    break
            else:
                search(idx+len(compare[i]))
    return
target= input().strip()
compare = []
dp = [0] * 101 # max 100+1
for _ in range(int(input())):
    compare.append(input().strip())
result = 0
search(0)
print(result)