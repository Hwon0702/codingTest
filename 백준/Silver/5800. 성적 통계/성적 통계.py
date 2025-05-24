import sys 
input = sys.stdin.readline
sys.setrecursionlimit(10**6)
n = int(input()) 
student = []
for j in range(1,n+1):
    student = list(map(int,input().split()))
    student = student[1:]
    student = sorted(student)
    lg = 0
    for i in range(0,len(student)-1): 
        if student[i+1] - student[i]>lg:
            lg = student[i+1] - student[i]
    print(f'Class {j}')
    print(f'Max {max(student)}, Min {min(student)}, Largest gap {lg}')