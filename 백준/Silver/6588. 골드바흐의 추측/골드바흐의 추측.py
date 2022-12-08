import sys 
input = sys.stdin.readline
sys.setrecursionlimit(10**6)
prime = [True for _ in range(1000001)]
prime[0]=False
prime[1]=False
j=2
for i in range(2, 1000001):
    if prime[i]:
        j=2
        while j*i<1000001:
            prime[i*j]=False
            j+=1

while True:
    n = int(input())
    if n==0:
        break
    else:
        flag = True
        for i in range(3,n):
            if prime[i] and prime[n-i] == True :
                print("%d = %d + %d"%(n , i , n-i))
                flag = False
                break
        if flag:
            print("Goldbach's conjecture is wrong.")
    