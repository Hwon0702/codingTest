import sys 
input = sys.stdin.readline 
sys.setrecursionlimit(10**6)
def get_gcd(x, y):
	while x > 0:
		x, y = y%x, x
	
	return y

a, b = map(int, input().split())
c, d = map(int, input().split())
gcd=get_gcd(b, d)
gcm = b*d/gcd 
a = a*(gcm/b)
c = c*(gcm/d)
new_gcd = get_gcd(a+c, gcm)
print(int((a+c)/new_gcd), int(gcm/new_gcd))