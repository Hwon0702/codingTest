import sys 
input = sys.stdin.readline 
sys.setrecursionlimit(10**6)
n, m = map(int, input().split())
books = list(map(int, input().split()))
pos = []
neg = []
count = []
for i in books:
    if i>0:
        pos.append(i)
    else:
        neg.append(abs(i))

pos.sort(reverse=True)
neg.sort(reverse=True)
#한 번에 가지고 움직일 수 있는 만큼
for i in range(len(pos)//m):
    count.append(pos[i*m]) #걸음 수 = 한 번에 움직인 수 중 더 큰 수
if len(pos)%m>0:#나누어 떨어지지 않을 경우
    count.append(pos[(len(pos)//m)*m])

for i in range(len(neg)//m):
    count.append(neg[i*m])
if len(neg)%m>0:
    count.append(neg[(len(neg)//m)*m])

count.sort()

res = count.pop()#종료했을 때
res += 2*sum(count)#왕복

print(res)