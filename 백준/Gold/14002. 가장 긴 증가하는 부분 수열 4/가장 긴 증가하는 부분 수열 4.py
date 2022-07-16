import sys 
input = sys.stdin.readline
sys.setrecursionlimit(10**6)
n = int(input())
res = [1 for _ in range(n)]
numbers = list(map(int, input().split()))
for i in range(n):
    for j in range(i):
        if numbers[i]>numbers[j] :
            res[i]=max(res[i], res[j]+1)
max_length= max(res)
print(max_length)
idx = res.index(max_length)
res_list = []
while idx>=0:
    if res[idx]==max_length:
        res_list.append(numbers[idx])
        max_length-=1
    idx -=1

for num in res_list[::-1]:
    print(num, end=' ')