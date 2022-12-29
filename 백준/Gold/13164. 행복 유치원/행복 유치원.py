import sys 
input = sys.stdin.readline
sys.setrecursionlimit(10**6)
#가장 키가 큰 원생과 가장 키가 작은 원생의 키 차이만큼
#K개의 조에 대해 티셔츠 만드는 비용의 합을 최소
n, k = map(int, input().split())
person = list(map(int, input().split()))
diffs = []
for i in range(1, n):
    diffs.append(person[i]-person[i-1])
diffs.sort(reverse=True)
if k==n:
    print(0)
else:
    print(sum(diffs[k-1:]))