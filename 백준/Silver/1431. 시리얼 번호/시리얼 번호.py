import sys 
input = sys.stdin.readline
sys.setrecursionlimit(10**6)
n = int(input())

arr = []
for i in range(n):
    arr.append(input().rstrip('\n'))


for i in range(n-1):
    for j in range(i+1, n):
        if len(arr[i]) > len(arr[j]):
            arr[i], arr[j] = arr[j], arr[i]

        elif len(arr[i]) == len(arr[j]):
            sum_a=0
            sum_b=0
            for x,y in zip(arr[i],arr[j]):
                if x.isdigit():
                    sum_a+=int(x)
                if y.isdigit():
                    sum_b+=int(y)
            if sum_a > sum_b:
                arr[i],arr[j] = arr[j], arr[i]

            elif sum_a == sum_b:
                for x,y in zip(arr[i], arr[j]):
                    if x > y:
                        arr[i],arr[j] = arr[j], arr[i]
                        break
                    elif x < y:
                        break
for v in arr:
    print(v)