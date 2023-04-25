import sys 
input = sys.stdin.readline
sys.setrecursionlimit(10**6)
n = int(input())
s = input().rstrip('\n')
res = 0
for i in range(n):
    res += (ord(s[i])-96) * (31 ** i)
print(res % 1234567891)