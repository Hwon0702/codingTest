import sys 
input = sys.stdin.readline
sys.setrecursionlimit(10**6)
tc = int(input())
for _ in range(tc):
    n = int(input())
    numbers = {}
    while(n!=1):
        if n%2==0:
            if numbers.get(2) is None:
                numbers[2]=1
            else:
                numbers[2]+=1
            n=n/2
        else:
            i=3
            while(n%i!=0):
                i+=1
            n=n/i
            if numbers.get(i) is None:
                numbers[i]=1
            else:
                numbers[i]+=1
    sorted_dict = sorted(numbers.items())
    for key in numbers:
        print(f"{key} {numbers[key]}")

