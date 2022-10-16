import sys 
input = sys.stdin.readline
sys.setrecursionlimit(10**6)
n = int(input())
graph = {}
for _ in range(n):
    root, left, right = map(str, input().split())
    graph[root]=[left, right]

def PreOrder(root):
    if root!='.':
        print(root, end='')
        PreOrder(graph[root][0])
        PreOrder(graph[root][1])
def InOrder(root):
    if root!='.':
        InOrder(graph[root][0])
        print(root, end='')
        InOrder(graph[root][1])
def PostOrder(root):
    if root!='.':
        PostOrder(graph[root][0])
        PostOrder(graph[root][1])
        print(root, end='')
PreOrder('A')
print()
InOrder('A')
print()
PostOrder('A')