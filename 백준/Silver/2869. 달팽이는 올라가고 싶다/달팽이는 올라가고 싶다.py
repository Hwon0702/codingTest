import math
import sys 
input = sys.stdin.readline
sys.setrecursionlimit(10**6)
A,B,V = map(int, input().split())
day_count =math.ceil((V-A)/(A-B)) + 1
print(day_count)