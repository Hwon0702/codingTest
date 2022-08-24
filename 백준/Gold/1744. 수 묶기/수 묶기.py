import sys
input = sys.stdin.readline
sys.setrecursionlimit(10**6)
n = int(input())
negative = []
positive = []
sum = 0
for _ in range(n):
    num =int(input())
    if num<=0:
        negative.append(num)
    elif num==1:
        sum+=1
    else:
        positive.append(num)
positive.sort()
negative.sort()
if len(positive)>0:
    if len(positive)%2==0:
        for i in range(0, len(positive), 2):
            sum += positive[i]*positive[i+1]
    else:
        sum += positive[0]
        for i in range(1, len(positive), 2):
            sum += positive[i]*positive[i+1]
if len(negative)>0:
    if len(negative)%2==0:
        for i in range(0, len(negative), 2):
            sum += negative[i]*negative[i+1]
    else:
        sum += negative[len(negative)-1]
        for i in range(0, len(negative)-1, 2):
            sum += negative[i]*negative[i+1]
print(sum)