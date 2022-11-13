import sys 
input = sys.stdin.readline
sys.setrecursionlimit(10**6)
n = int(input())
numbers = list(map(int, input().split()))
res = []

def find(start, end, target):
    while start<end:
        mid = (start+end)//2
        if res[mid]<target:
            start = mid+1
        else:
            end=mid
    return end

for n in numbers:
    if len(res)<=0:
        res.append(n)
        continue
    if res[-1]<n:
        res.append(n)
    else:
        idx = find(0, len(res)-1, n)
        res[idx]=n
print(len(res))