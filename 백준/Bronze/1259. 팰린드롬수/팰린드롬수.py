import sys 
input = sys.stdin.readline
sys.setrecursionlimit(10**6)

def palindrome(s):
    s_arr = list(s)
    for i in range(len(s_arr)//2):
        if s_arr[i] != s_arr[len(s_arr)-i-1]:
            return "no"
    return "yes"
        
while True:
    s = input().strip('\n')
    if s=='0':
        break
    print(palindrome(s))