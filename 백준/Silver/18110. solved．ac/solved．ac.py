import sys 
input = sys.stdin.readline
sys.setrecursionlimit(10**6)
n= int(input())
def custom_round(num):
    if num - int(num) >= 0.5:
        return int(num)+1
    else:
        return int(num)
 
if n==0:
    print(0)
else:
    scores = []
    for _ in range(n):
        scores.append(int(input()))
    scores.sort()
    remove_cnt = custom_round(n*0.15)
    if remove_cnt!=0:
        scores = scores[remove_cnt:n-remove_cnt]
    print(custom_round(sum(scores)/(n-(remove_cnt*2))))