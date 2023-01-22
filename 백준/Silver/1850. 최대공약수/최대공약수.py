import sys 
input = sys.stdin.readline
sys.setrecursionlimit(10**6)
def gcd(a, b):
	while a > 0 :
		a, b = b%a, a	
	return b

a, b = map(int, input().split())
res = gcd(a,b)
for _ in range(res):
	print(1, end='')
