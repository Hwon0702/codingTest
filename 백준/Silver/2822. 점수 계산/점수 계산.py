#참가자의 총 점수는 가장 높은 점수 5개의 합
score = [0]*8
sum  = 0
sum =0
get_score=[]
get = []
for i in range(8):
    score = int(input())
    get_score.append((score,i+1))
get_score.sort()
for i in range(5):
    sum += get_score[::-1][i][0]
    get.append(get_score[::-1][i][1])
get.sort()
print(sum)
print(*get)