from operator import contains
import sys 
input = sys.stdin.readline
sys.setrecursionlimit(10**6)
a, b = map(str, input().split())
min = int(a.replace('6', '5')) + int(b.replace('6', '5'))
max = int(a.replace('5', '6')) + int(b.replace('5', '6'))
print(min, max)