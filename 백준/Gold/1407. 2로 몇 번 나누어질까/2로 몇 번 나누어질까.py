import sys 
input = sys.stdin.readline
sys.setrecursionlimit(10**6)
a, b=map(int, input().split())
def cal(num):
    cnt = num
    i = 2
    while i <= num:
        cnt += (num//i)*(i//2)
        i *= 2
    return cnt
print(cal(b)-cal(a-1))