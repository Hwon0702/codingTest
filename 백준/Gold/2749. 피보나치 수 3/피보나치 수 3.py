import sys 
input = sys.stdin.readline
sys.setrecursionlimit(10**6)
M=1000000
P=1500000
n = int(input())
a, b = 0, 1
for i in range(n%P):
    a, b = b%M, (a+b)%M
    
print(a) 
