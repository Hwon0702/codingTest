import sys 
from collections import deque
input = sys.stdin.readline
sys.setrecursionlimit(10**6)
n = int(input())
queue = deque()
queue.append((1, 0))
already = {}
already[(1, 0)]=0
while queue:
    emoji_cnt, clipboard = queue.popleft()
    if emoji_cnt==n:
        print(already[(emoji_cnt, clipboard)])
        break
    else:
        #화면에 있는거 모두 복사
        if (emoji_cnt, emoji_cnt) not in already.keys():
            already[(emoji_cnt, emoji_cnt)]=already[(emoji_cnt, clipboard)]+1
            queue.append((emoji_cnt, emoji_cnt))
        #클립보드 붙여넣기
        if (emoji_cnt+clipboard, clipboard) not in already.keys():
            already[(emoji_cnt+clipboard, clipboard)]=already[(emoji_cnt, clipboard)]+1
            queue.append((emoji_cnt+clipboard, clipboard))
        #화면에 있는거 하나 지우기
        if (emoji_cnt-1, clipboard) not in already.keys():
            already[(emoji_cnt-1, clipboard)]=already[(emoji_cnt, clipboard)]+1
            queue.append((emoji_cnt-1, clipboard))