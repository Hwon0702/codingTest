import sys 
input=sys.stdin.readline
sys.setrecursionlimit(10**6)
s = input().strip('\n')
s_arr = list(s)
cnt = 0
for i in range(1, len(s)):
    if s_arr[i-1]!=s_arr[i]:
        cnt+=1
print((cnt+1)//2)