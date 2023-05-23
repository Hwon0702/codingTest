import sys 
input = sys.stdin.readline 
sys.setrecursionlimit(10**6)

n = int(input())
if n%2==0:
    print("SK")
else:
    print("CY")