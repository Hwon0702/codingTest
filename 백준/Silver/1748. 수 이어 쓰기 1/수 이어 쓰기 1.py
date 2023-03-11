import sys 
input = sys.stdin.readline
sys.setrecursionlimit(10**6)
n = int(input())
length = len(str(n)) - 1
res = 0
for i in range(length):
    res += 9 * (10 ** i) * (i + 1)
    i += 1
res += ((n - (10 ** length)) + 1) * (length + 1)

print(res)