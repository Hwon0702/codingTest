import sys 
input = sys.stdin.readline 
sys.setrecursionlimit(10**6)
n = int(input())
methods = list(map(str, input().rstrip('\n')))
num_list = []
for _ in range(n):
    num_list.append(int(input()))
stack = []

for i in methods :					
    if 'A' <= i <= 'Z' :
        stack.append(num_list[ord(i)-ord('A')])
    else :
        s2 = stack.pop()
        s1 = stack.pop()

        if i =='+' :
            stack.append(s1+s2)
        elif i == '-' :
            stack.append(s1-s2)
        elif i == '*' :
            stack.append(s1*s2)
        elif i == '/' :
            stack.append(s1/s2)

print('%.2f' %stack[0])