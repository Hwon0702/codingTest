import sys 
import math
input = sys.stdin.readline 
sys.setrecursionlimit(10**6)


def isPrime(x):
    if x == 1:
        return False
    for i in range(2, int(math.sqrt(x)+1)):
        if x % i == 0:
            return False
    return True

def isPalindrom(x):
    x = str(x)
    tmp = x[::-1]
    if x==tmp:
        return True
    else:
        return False

n = int(input())

while True:
    if isPalindrom(n):
        if isPrime(n):
            print(n)
            break

    n+=1