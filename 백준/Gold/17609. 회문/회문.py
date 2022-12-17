import sys 
input = sys.stdin.readline
sys.setrecursionlimit(10**6)
n = int(input())
def similar_palindrome(target, s, e):
    while (s<e):
        if target[s]==target[e]:
            s+=1
            e-=1
        else:
            return False
    return True

def palindrome(target, s, e):
    while s<e:
        if target[s]==target[e]:
            s+=1
            e-=1
        else :
            rm_left = similar_palindrome(target, s+1, e)
            rm_right = similar_palindrome(target, s, e-1)
            if rm_left or rm_right:
                return 1
            else:
                return 2
    return 0
for _ in range(n):
    words = list(map(str, input().strip('\n')))
    print(palindrome(words, 0, len(words)-1))