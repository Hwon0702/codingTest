import sys 
input = sys.stdin.readline
sys.setrecursionlimit(10**6)
from itertools import combinations, product,chain
vowel_list=["a","e", "i", "o", "u"]
consonant_target =[]
vowel_target=[]
l, c = map(int, input().split())
str_list = list(map(str, input().split()))
for s in str_list :
    if s in vowel_list:
        vowel_target.append(s)
    else:
        consonant_target.append(s)
if len(str_list)>0:
    res_result = []
    for i in range(1, len(consonant_target)+1):
        if l-i>=2:
            vowel_comb = list(combinations(vowel_target, i))
            consonant_comb = list(combinations(consonant_target, l-i))
            vowel_comb= list(map(list, vowel_comb))
            consonant_comb= list(map(list, consonant_comb))
            product_list = list(product(vowel_comb, consonant_comb))
            for v in product_list:
                res_list = list(chain(*v))
                res_list.sort()
                res_result.append(''.join(map(str, res_list)))
    res_result = list(set(res_result))
    res_result.sort()
    for v in res_result:
        print(v)


