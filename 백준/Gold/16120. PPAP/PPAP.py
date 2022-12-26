import sys 
input = sys.stdin.readline 
sys.setrecursionlimit(10**6)
target = map(str, input().strip('\n'))
if target =='PPAP' or target =='P':
    print("PPAP")
else:
    ppap = ['P', 'P', 'A', 'P']
    s = []
    
    for i in target:
        s.append(i)
        if s[-4:]==ppap:
            s.pop()
            s.pop()
            s.pop()
    if s ==ppap or s ==['P']:
        print('PPAP')
    else:
        print('NP')