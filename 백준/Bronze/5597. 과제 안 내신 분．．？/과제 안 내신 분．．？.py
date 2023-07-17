import sys 
input = sys.stdin.readline 
sys.setrecursionlimit(10**6)
person=[False for _ in range(31)]
for _ in range(28):
    n = int(input())
    person[n]=True
for i in range(1, 31):
    if person[i]==False:
        print(i)