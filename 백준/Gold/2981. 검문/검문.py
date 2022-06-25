from math import sqrt, gcd
import sys 
input = sys.stdin.readline
n = int(input())
diff = []
get_gcd = 0
numbers = [0]*n
results = []
for i in range(0, n):
    numbers[i] = int(input())
numbers.sort()
for i in range(1, n):
    diff.append(numbers[i]-numbers[i-1])
get_gcd = diff[0]
for i in range(1, len(diff)):
    get_gcd = gcd(get_gcd, diff[i])

for i in range(2, int(sqrt(get_gcd)) + 1): #제곱근까지만 탐색
    if get_gcd % i == 0:
        results.append(i)
        results.append(get_gcd//i)
results.append(get_gcd)
results = list(set(results)) 
results.sort()
for v in results:
    print(v, end=' ')
