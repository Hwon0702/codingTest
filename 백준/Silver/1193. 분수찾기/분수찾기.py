import sys 
input = sys.stdin.readline 
sys.setrecursionlimit(10**6)
n = int(input())
line = 0
end = 0
while n > end:
    line += 1
    end += line

diff = end - n
if line%2 == 0: #짝수 
    top = line - diff
    bottom = diff + 1
else:
    top = diff + 1
    bottom = line - diff

print("%d/%d"%(top,bottom))