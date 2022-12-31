import sys 
input = sys.stdin.readline 
sys.setrecursionlimit(10**6)
a, b = map(str, input().split())
diff = []
for i in range(len(b)-len(a)+1):
    count = 0
    for j in range(len(a)):
        if a[j]!=b[i+j]:
            count+=1
    diff.append(count)
print(min(diff))