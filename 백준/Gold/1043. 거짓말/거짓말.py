import sys 
input = sys.stdin.readline
sys.setrecursionlimit(10**6)

n, m = map(int, input().split())
t = list(map(int, input().split()))
t = set(t[1:]) #진실을 아는 사람 set
graph = [{}for _ in range(m)] #파티 갯수만큼 미리 만들어두기
party = [False] * m #진실만 말해야 하는 파티
for i in range(m):
    v = list(map(int, input().split()))
    v =v[1:]
    graph[i] =set(v) #파티별 참석자
for _ in range(m): #각 파티에 참석한 사람 확인
    for i in range(m):
        if t & graph[i]: #파티 참석자중 진실을 아는 사람 섞여있음(=교집합 0 아님)
            party[i]=True #진실만 말해야함
            t = t | graph[i] #진실을 아는 사람 추가
print(party.count(False)) #거짓 말해도 되는 갯수 확인