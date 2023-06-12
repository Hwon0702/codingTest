import sys 
input = sys.stdin.readline 
sys.setrecursionlimit(10**6)
n = int(input())
a = list(map(int, input().split()))
b = list(map(int, input().split()))

a.sort()
b.sort(reverse=True) #가장 차이가 적게 이기게 하기 위함
#경기여부 표시
a_visited = [False for _ in range(n)]
b_visited = [False for _ in range(n)]

score = 0

#a가 이김
for i in range(n):
    for j in range(n):
        if a[i] > b[j] and a_visited[i]==False and b_visited[j]==False:
            score += 2
            a_visited[i]=True
            b_visited[j]=True
            break
#비김
for i in range(n):
    if a_visited[i] == True:
        continue
    for j in range(n):
        if a[i] == b[j]  and a_visited[i]==False and b_visited[j]==False:
            score += 1
            a_visited[i]=True
            b_visited[j]=True
           
print(score)