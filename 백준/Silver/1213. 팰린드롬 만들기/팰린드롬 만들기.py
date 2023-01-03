import sys 
input = sys.stdin.readline
sys.setrecursionlimit(10**6)
words = input().strip('\n')
alphabets = [0 for _ in range(27)]  
odds = 0                               
oddIndex = -1
left=[]
for letters in words:
    alphabets[ord(letters)-65] += 1
for i in range(len(alphabets)):
    if alphabets[i] % 2 == 1:              
        odds += 1
        oddIndex = i    
        if odds >= 2:
            print("I'm Sorry Hansoo")
            exit()
alphabets[oddIndex] -= 1 #홀수일 경우 미리 빼놓기
for i in range(len(alphabets)): 
    if alphabets[i] % 2 != 1:
        for _ in range(alphabets[i]//2):
            left.append(chr(i+65))
            print(chr(i+65), end="")
if oddIndex != -1:
    print(chr(oddIndex+65), end="") 
left.reverse()
for v in left: #오른쪽 = 왼쪽을 뒤집은거
    print(v, end="")