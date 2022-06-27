import sys
input = sys.stdin.readline
sys.setrecursionlimit(10**6)
# 자연수 A를 B번 곱한 수를 알고 싶다. 단 구하려는 수가 매우 커질 수 있으므로 이를 C로 나눈 나머지를 구하는 프로그램을 작성하시오.
#A**b %c
def get_power(a, b, c):
    if b == 0:
        return 1 #A의 0승은 1
    elif b == 1:
        return a%c #A*1 = A
    res = get_power(a, b//2, c) % c ##그냥 곱(제곱)이므로, b //2를 하면 제곱근으로 나눈것과 동일(지수법칙 : a^(n+m) = a^n * a^m)
    if b % 2 == 0: # (a*b)%c = (a%c * b%c)%c
        return (res % c) * (res % c) % c	
    return (res % c) * (res % c) * a % c

a, b, c = list(map(int, input().split()))
print(get_power(a, b, c))