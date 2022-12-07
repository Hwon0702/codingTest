import sys 
import math
input = sys.stdin.readline
sys.setrecursionlimit(10**6)
n = int(input())
if n<8:
    print(-1)
else:
    nums = [True for _ in range(n+2)]
    nums[0]=False
    nums[1]=False
    primes =[]
    for i in range(2, int(math.sqrt(n)) + 1): 
        if nums[i]==True:
            j=2
            while i*j<n:
                nums[i*j]=False
                j+=1
    #print(nums)
    if n%2==0:
        ans=[2,2]
        n-=4
        for i in range(2, n+1):
            if nums[i] and nums[n-i]:
                ans.append(i)
                ans.append(n-i)
                break
        print(*ans)
    else:
        ans=[2,3]
        n-=5
        for i in range(2, n+1):
            if nums[i] and nums[n-i]:
                ans.append(i)
                ans.append(n-i)
                break
        print(*ans)