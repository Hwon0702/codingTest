import sys 
input = sys.stdin.readline 
sys.setrecursionlimit(10**6)
#문자열 S의 길이는 N이고, 'A', 'B'로 이루어져 있다.
#문자열 S에는 0 ≤ i < j < N 이면서 s[i] == 'A' && s[j] == 'B'를 만족하는 (i, j) 쌍이 K개가 있다.

n,k = map(int,input().split())
a = 0
b = n
while a*b < k and b > 0:
    a += 1
    b -= 1
if k == 0:
    print('B'*n)
    sys.exit()
elif b == 0:
    print(-1)
    sys.exit()
remain = k - (a-1)*b
print('A'*(a-1) + 'B'*(b-remain) + 'A' + 'B'*remain)
