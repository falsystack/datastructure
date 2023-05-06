# Array
- 배열, 연속된 메모르 블럭
  - 한번에 할당되고 한번에 해제된다.
- Random Access 에 강하다
  - `O(1)`이다 
  - ex: `arr[33] = 100`
- 삽입/삭제에 약하다
  - `O(N)`
- Cache Miss가 잘 일어나지 않는다.
  - Cache: CPU에 달려있는 캐시(Cache)메모리
- 요소가 사라질 떄 마다 GC되지 않는다.
  - GC는 GC대상의 메모리 사이즈보다 GC 대상 갯수에 상관이 더 많다.

# Linked List
- 포인터로 연결한 노드들로 구성된 자료구조
- go에서는 container라는 패키지에서 제공되고 있다.
- 불연속 메모리다.
  - 필요한 만큼만 메모리 사용한다.
- 삽입 / 삭제에 효율적
  - 포인터만 바꿔주면 된다.
- `Graph`의 일종이다.
  - `Graph`는 node - edge로 이루어진 자료구조이다.
  - Linked List는 엣지가 하나인 것을 말한다.
- Single Linked List
  - 엣지가 단방향
- Double Linked List(go/container)
  - 엣지가 양방향
- Cache Miss 가 잘일어남. (고도의 밀집화된 데이터의 처리를 할 때 굉장히 비효율적이다.)
- 요소가 사라질 때 GC가 일어남

# 결론
대부분 배열(Array)를 사용하는것이 좋다.
- 단 요소수가 많고 맨 앞에서 삽입 / 삭제가 빈번한 큐는 Linked List사용 고려
  - 사용예: 대기열(FIFO)
- 멀티 쓰레드 환경에서 Lock free Queue 등도 고려