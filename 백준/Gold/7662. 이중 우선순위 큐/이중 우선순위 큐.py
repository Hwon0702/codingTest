import sys 
input = sys.stdin.readline 
sys.setrecursionlimit(10**6)
import heapq

def calculate(case):
    asc_queue = []
    desc_queue = []
    removed = [False for _ in range(case+1)]
    for idx in range(case):
        m, n = input().split()
        n = int(n)
        if m=='I':
            heapq.heappush(asc_queue, (n, idx))
            heapq.heappush(desc_queue,(-1*n, idx))
        else:
            if n<0:
                while asc_queue:
                    if removed[asc_queue[0][1]]==True :
                        heapq.heappop(asc_queue)
                        removed[r]=True
                    else:
                        _, r = heapq.heappop(asc_queue)
                        removed[r]=True
                        break
            else:
                while desc_queue:
                    if removed[desc_queue[0][1]]==True :
                        heapq.heappop(desc_queue)
                    else:
                        _, r = heapq.heappop(desc_queue)
                        removed[r]=True
                        break
                    
    asc = []
    desc = []

    while asc_queue:
        if removed[asc_queue[0][1]]==True :
            heapq.heappop(asc_queue)
        else:
            n, r = heapq.heappop(asc_queue)
            asc.append(n)
    while  desc_queue:
        if removed[desc_queue[0][1]]==True :
            heapq.heappop(desc_queue)
        else:
            n, r = heapq.heappop(desc_queue)
            desc.append(n*(-1))
    if len(asc)==0 and len(desc)==0:
        print("EMPTY")
    else:
        if len(asc)> 0 and len(desc)>0:
            print(desc[0], asc[0])
        elif len(asc)>0 and len(desc)==0:
            print(asc[0], asc[0])
        else:
            print(desc[0], desc[0])

tc = int(input())
for _ in range(tc):
    calculate(int(input()))
