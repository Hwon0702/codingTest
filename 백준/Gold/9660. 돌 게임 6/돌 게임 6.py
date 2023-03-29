import sys 
input = sys.stdin.readline 
sys.setrecursionlimit(10**6)
n=int(input())
if n%7 == 2 or n%7 == 0:
    print("CY")
else:
    print("SK")