import sys 
input = sys.stdin.readline 
sys.setrecursionlimit(10**6)
tc = int(input())
is_prime = [True for _ in range(1000001)] 
is_prime[0]=False
is_prime[1]=False
for i in range(2, 1000001):
    if is_prime[i]==True:
        for j in range(2, 1000001):
            if i*j>=1000001:
                break
            else:
                is_prime[i*j]=False
def find(n):
    cnt = 0
    for i in range(2,n//2+1):
        if is_prime[i] and is_prime[n-i]:
            cnt += 1
    print(cnt)

for _ in range(tc):
    n = int(input())
    find(n)
