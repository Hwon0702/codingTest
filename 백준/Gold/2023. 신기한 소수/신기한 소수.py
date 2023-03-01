import sys 
input = sys.stdin.readline 
sys.setrecursionlimit(10**6)

n = int(input())
primeSet = set()
def isPrime(num):
    if num<2:
        return False
    if num in primeSet:
        return True
    for i in range(2,int(num**0.5)+1):
        if(num%i==0):
            return False
    primeSet.add(num)
    return True

def dfs(num):
  if len(str(num)) == n:
    print(num)
  else:
    for i in range(1,10,2):
      if isPrime(10 * num + i):
        dfs(10 * num + i)
      else:
        continue

#일단 끝자리가 2,3,5,7이어야함
dfs(2)
dfs(3)
dfs(5)
dfs(7)