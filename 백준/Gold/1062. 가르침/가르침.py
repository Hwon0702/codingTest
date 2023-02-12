import sys 
input = sys.stdin.readline 
#a c i n t 
n, k = map(int, input().split())
alphabets = ['a', 'n', 't', 'i', 'c']
alphabets_list = ['b', 'd', 'e', 'f', 'g', 'h', 'j', 'k', 'l', 'm','o', 'p', 'q', 'r', 's', 'u', 'v', 'w', 'x', 'y', 'z']
words = []
result =0 
def select(n, start):
    global result
    if n == 0:
        result = max(result, check())
        return
    for i in range(start, len(alphabets_list)):
        alphabets.append(alphabets_list[i])
        select(n-1, i+1)
        alphabets.pop()

def check():
    result = 0
    for w in words:
        isRead = True
        for i in range(4, len(w)-4):
            if w[i] not in alphabets:
                isRead = False
                break
        if isRead:
            result += 1
    return result

for _ in range(n):
    words.append(input().strip('\n'))
if k<5:
    print(0)
elif k>=26:
    print(n)
else:
    select(k-5, 0)
    print(result)