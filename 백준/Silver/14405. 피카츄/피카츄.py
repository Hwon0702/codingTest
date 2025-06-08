import sys 
input = sys.stdin.readline
sys.setrecursionlimit(10**6)
s = str(input().rstrip())
s = str.replace(s, "pi", "1")
s = str.replace(s, "ka", "1")
s = str.replace(s, "chu", "1")
s = str.replace(s, "1", "")
if len(s)>=1:
    print("NO")
else:
    print("YES")