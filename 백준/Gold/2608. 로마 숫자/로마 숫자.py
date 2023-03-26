from collections import deque
import sys 
input = sys.stdin.readline 
sys.setrecursionlimit(10**6)
roman1 = list(map(str, input().strip('\n')))
roman2 = list(map(str, input().strip('\n')))

roman_num = {'I':1, 'IV':4, 'V':5, 'IX':9, 'X':10, 'XL':40, 'L':50,'XC':90, 'C':100, 'D':500, 'CD':400,'CM':900,'M':1000}
num_roman = {1:'I', 4:'IV', 5:'V', 9:'IX', 10:'X', 40:'XL', 50:'L',90:'XC', 100:'C', 500:'D', 400:'CD',900:'CM',1000:'M'}
def convert_num(roman):
    num = []
    while roman:
        l = roman.pop()
        cnt = 1
        if len(roman)>0 and roman[len(roman)-1]+l in roman_num:
            l=roman.pop()+l
            num.append(roman_num[l]*cnt)
            continue
        
        while len(roman)>0 and l==roman[len(roman)-1]:
            cnt+=1
            roman.pop()
        num.append(roman_num[l]*cnt)
    return sum(num)
def convert_roman(num):
    roman = []
    num_list = list(str(num))
    idx = len(num_list)-1
    lst = ['I', 'X', 'C', 'M']
    lst2 = ['V', 'L', 'D']    
    lst3 = ['IV', 'XL', 'CD'] 
    lst4 = ['IX', 'XC', 'CM'] 
    while num_list:
        c = num_list.pop(0)
        if c == '4':  
            roman += lst3[idx]
        elif c == '9':
            roman += lst4[idx]
        elif c in '123': 
            for _ in range(int(c)):
                roman += lst[idx]
        elif c in '5678': 
            roman += lst2[idx]
            if c != '5':
                for _ in range(int(c)-5):
                    roman += lst[idx]
        idx -= 1
    return roman

num1 = convert_num(roman1)
num2 = convert_num(roman2)
roman = convert_roman(num1+num2)
print(num1+num2)
print(''.join(roman))
    
