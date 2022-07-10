import sys
input = sys.stdin.readline
sys.setrecursionlimit(10**6)
tc = int(input())

for i in range(0, tc):
    cnt = 1
    people = []
    n = int(input())
    for i in range(n):
        paper_score, interview_score = map(int,input().split())
        people.append([paper_score, interview_score])
    people.sort() # 서류 기준 오름차순
    max = people[0][1]
    
    for i in range(1,n):
        if max > people[i][1]:
            cnt += 1
            max = people[i][1]
    print(cnt)
