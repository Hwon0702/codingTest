import sys
input=sys.stdin.readline
sys.setrecursionlimit(10**6)
n=int(input())
numbers=[]
start=[0,0,0,0,0,0,0,0,0,0]
alphabets = {}
for i in range(65, 75):
    alphabets[chr(i)]=i
for _ in range(n):
    a=input().strip('\n')
    numbers.append(a)
    start[alphabets[a[0]]-65]=1
    
cnt=[[0, x] for x in range(10)]

for i in range(n):
    p = 1
    for j in range(len(numbers[i])-1, -1, -1):
        cnt[alphabets[numbers[i][j]]-65][0]+=pow(10, p)
        p+=1
cnt.sort(reverse=True)

dic=[0,0,0,0,0,0,0,0,0,0]
for i in range(10):
    dic[9-i]=chr(cnt[i][1]+65)
    if i==9 and start[cnt[i][1]]==1:
        for j in range(1, 10):
            if start[alphabets[dic[j]]-65]==0:
                t=dic.pop(j)
                break
        dic.insert(0, t)

res = 0
for i in range(n):
    new=[]
    for j in numbers[i]:
        new.append(str(dic.index(j)))
    res+=int(''.join(new))

print(res)