import sys 
input = sys.stdin.readline
sys.setrecursionlimit(10**6)
n= int(input())
list = [0] * 26

for _ in range(n) :
    name = input()
    list[ord(name[0])-97] += 1
if max(list) >= 5 :
    for i in range(len(list)) :
        if list[i] >= 5 :
            print(chr(i+97), end='')
else :
    print("PREDAJA")