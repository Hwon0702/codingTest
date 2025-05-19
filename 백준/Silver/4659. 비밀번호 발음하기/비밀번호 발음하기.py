import sys 
input = sys.stdin.readline
sys.setrecursionlimit(10**6)
vowel = ['a','e','i','o','u']

while True:
    result = True
    word = input().rstrip('\n')
    if word=='end':
        break
    words = list(word)
    for i in range(1, len(words)):
        if words[i]==words[i-1] and words[i] not in ['e', 'o']:
            result=False
    
    vowel_cnt = 0
    consonant_cnt = 0
    total_vowel_cnt = 0
    for w in words:
        if w in vowel:
            vowel_cnt+=1
            consonant_cnt=0
            total_vowel_cnt+=1
        else:
            consonant_cnt+=1
            vowel_cnt=0
        if vowel_cnt>=3 or consonant_cnt>=3:
            result=False
    if total_vowel_cnt<=0:
        result=False
    if result:
        print("<{0}> is acceptable.".format(word))
    else:
        print("<{0}> is not acceptable.".format(word))