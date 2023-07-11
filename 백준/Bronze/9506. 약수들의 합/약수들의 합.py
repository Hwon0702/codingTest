import sys 
input = sys.stdin.readline
sys.setrecursionlimit(10**6)
while True:
    n = int(input())
    if n==-1:
        break
    nums = set()

    for i in range(2, int(n**0.5)+1):
        if n%i==0:
            nums.add(i)
            nums.add(int(n/i))
    nums.add(1)
    num= list(nums)
    num.sort()
    if sum(num)==n:
        print("{} = ".format(n), end='')
        for v in range(len(num)):
            if v==len(num)-1:
                print(num[v])
            else:
                print("{} + ".format(num[v]), end='')
    else:
        print("{} is NOT perfect.".format(n))