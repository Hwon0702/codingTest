import sys 
input = sys.stdin.readline
sys.setrecursionlimit(10**6)
#문자열 뒤에 A
#B추가 후 뒤집기

def dfs(target):
    global flag
    if len(S)==len(target):
        if S == target:
            flag=1
        return
    if target[-1]=='A':
        target.pop()
        dfs(target)
        target.append('A')

    if target[0]=='B':
        target.reverse()
        target.pop()
        dfs(target)
        target.append('B')
        target.reverse()
S = list(map(str, input().strip('\n')))
T = list(map(str, input().strip('\n')))
flag = 0
dfs(T)

print(flag)
