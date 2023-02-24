import sys 
input = sys.stdin.readline 
sys.setrecursionlimit(10**6)
numbers = list(map(str, input().rstrip('\n')))
cnts = [0 for _ in range(9)]
for i in numbers:
    if i=='6' or i =='9':
        cnts[6]+=1
    else:
        cnts[int(i)]+=1
if cnts[6] % 2 == 0:
    cnts[6] = cnts[6] // 2
else:
    cnts[6] = (cnts[6] // 2)+1
print(max(cnts))