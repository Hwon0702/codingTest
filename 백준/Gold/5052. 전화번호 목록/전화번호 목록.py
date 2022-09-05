import sys 
input = sys.stdin.readline
sys.setrecursionlimit(10**6)

def cal(n):
    strs = []
    for _ in range(n):
        strs.append(input().strip('\n'))
    strs.sort()
    ret = "YES"
    break_flag = False
    for i in range(0, n-1):
        if len(strs[i+1])>=len(strs[i]) and strs[i+1][:len(strs[i])]==strs[i]:
            ret = "NO"
            break
    print(ret)    
tc= int(input())
for i in range(tc):
    cal(int(input()))