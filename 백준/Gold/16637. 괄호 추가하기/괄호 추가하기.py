import sys
input = sys.stdin.readline
sys.setrecursionlimit(10**6)

def calc(num1, method, num2):
    if method == '+':
        return num1+num2
    elif method == '-':
        return num1-num2
    elif method == '*':
        return num1*num2
    
def dfs(idx, t):
    global res
    if idx == n-1:
        res = max(res, t)
        return
    if idx +2 <n:
        dfs(idx+2, calc(t,equal[idx+1], equal[idx+2]))
    if idx+4<n:
        dfs(idx+4, calc(t,equal[idx+1],(calc(equal[idx+2],equal[idx+3], equal[idx+4]))))
n = int(input())
equal = list(input().rstrip())
res = -float('inf')
for i in range(0,len(equal),2):
    equal[i] = int(equal[i])
t = equal[0]
dfs(0,t)
print(res)