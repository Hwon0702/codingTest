import sys
input = sys.stdin.readline 
sys.setrecursionlimit(10**6)
n1, n2 = map(int, input().split())

ant1 = list(map(str,input().rstrip('\n')))
ant2 = list(map(str,input().rstrip('\n')))
t = int(sys.stdin.readline())

ant1.reverse()
total = ant1 + ant2 
for _ in range(t):
    for i in range(len(total) - 1):
        if total[i] in ant1 and total[i + 1] in ant2:
            total[i], total[i + 1] = total[i + 1], total[i]
            if total[i + 1] == ant1[-1]:
                break

print(''.join(total))