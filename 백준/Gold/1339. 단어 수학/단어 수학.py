import sys 
input = sys.stdin.readline
sys.setrecursionlimit(10**6)
n = int(input())
#가장 자릿수가 큰 숫자가 : 가장 큰 수를 가져감
#자릿수가 동일하다면 : 많이나오는게 가져감
strings_list = ['' for _ in range(n)]
for i in range(n):
    strings_list[i] = input().strip()
dict = {}
for word in strings_list:
    cnt = len(word)-1
    for w in word:
        if w not in dict:
            dict[w] = pow(10, cnt)
        else:
            dict[w] += pow(10, cnt)
        cnt -= 1
dict = sorted(dict.values(), reverse=True)
result = 0
num = 9
for v in dict:
    result += v*num
    num -= 1
print(result)
